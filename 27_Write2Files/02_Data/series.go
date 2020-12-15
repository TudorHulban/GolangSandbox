package series

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"../../32_XLS/exportxls"
	simplify "../../42_LineSimplification/simplifybypoints"
	// "../../42_LineSimplification/simplify"
)

// InitSimplify Initialization params for the simplification operation.
type InitSimplify struct {
	ProcessID         string // ID given to instance to track / debug it
	ExportCSV         string // path to CSV file that contains selected positions
	ExportXLS         string // path to CSV file that contains selected positions
	SimplifyTolerance float64
	InputPositions    []int // holds positions of series to be considered. if bool the series needs to be simplified
	SimplifyPositions []int
}

// CfgCSV CSV configuration.
type CfgCSV struct {
	InitSimplify
	InputCSV string // path to CSV file that serves as input. TODO: add format of the input file

}

// CfgStream CSV configuration.
type CfgStream struct {
	InitSimplify
	InputStream [][]interface{} // data can be mixed this is why using interface
}

// Data structure holds series input and output information. To be used within a pool.
// Methods take a source stream or CSV and extract the series selected.
type Data struct {
	State        int
	BasePosition int             // Holds base position as part of input positions. work only with one base for now. Position starts from zero. Position zero is first position from input positions.
	ChoppedData  [][]interface{} // holds data after chop
	remappedData [][]interface{} // holds data after remap. TODO: what is remap?

	combinedData [][]simplify.Point // holds remapped data that was processed by combiner

	/*
	 Holds positions of output stream (and not original positions) that should be simplified. Position start from zero.
	 Position zero is first position from combined data.
	 Base series could also be simplified.
	*/

	SimplifiedData [][]interface{} // holds chopped data after simplification
	SimplifiedCSV  string          // path to CSV file that contains simplified series
	SimplifiedXLS  string          // path to CSV file that contains simplified series

	errorList []error // accumulated errors
}

// steps follows an operational flow. should be immutable.
func steps() map[int]string {
	state = make(map[int]string, 10)

	state[0] = "init"
	state[1] = "chop"
	state[2] = "remap"
	state[3] = "exportcsv"
	state[4] = "exportxls"
	state[5] = "combine"
	state[6] = "simplify"
	state[7] = "exportcsvsimple"
	state[8] = "exportxlssimple"
	state[9] = "park"

	return state
}

// NewCSV Creates instance based on CSV file. Positions start from 0.
func NewCSV(cfg CfgSimplify) (*Data, error) {
	if pBasePosition >= len(pPositions) {
		return nil, errors.New("base position outside selected series")
	}

	instance := new(Data)
	instance.ProcessID = "D" + strconv.FormatInt(pProcessID, 10)
	instance.InputCSV = pCSV
	instance.InputPositions = pPositions
	instance.SimplifyPositions = pSimplifyPositions
	instance.ExportCSV = pExportCSV
	instance.ExportXLS = pExportXLS
	instance.BasePosition = pBasePosition
	instance.SimplifyTolerance = pSimpleTolerance
	return instance, nil
}

/*
NewStream is constructor type. Creates instance based on stream. Positions start from 0.
*/
func NewStream(pProcessID int64, pStream [][]interface{}, pPositions, pSimplifyPositions []int, pExportCSV, pExportXLS string, pBasePosition int, pSimpleTolerance float64) (*Data, error) {
	if pBasePosition >= len(pPositions) {
		return nil, errors.New("base position outside selected series")
	}
	instance := new(Data)
	instance.ProcessID = "D" + strconv.FormatInt(pProcessID, 10)
	instance.InputStream = pStream
	instance.InputPositions = pPositions
	instance.SimplifyPositions = pSimplifyPositions
	instance.ExportCSV = pExportCSV
	instance.ExportXLS = pExportXLS
	instance.BasePosition = pBasePosition
	instance.SimplifyTolerance = pSimpleTolerance
	return instance, nil
}

// Init provides extended initialization to the instance.
func (d *Data) Init() {
	d.ChoppedData = make([][]interface{}, len(d.InputPositions))
	d.State = 0
}

// Auto drives the flow forward. Provides clear picture of where the processing is. Each step with logging.
func (d *Data) Auto() error {
	nextStep := func() {
		d.State++
		d.Auto()
	}

	switch d.State {
	case 0: // "init"
		{
			log.Println("state: ", d.ProcessID, state[d.State])
			d.Init()
			nextStep()
		}
	case 1: // "chop"
		{
			log.Println("state: ", state[d.State])
			if len(d.InputStream) == 0 {
				errChop := d.ChopCSV()
				if errChop != nil {
					d.errorList = append(d.errorList, errChop)
					return errChop
				}
			} else {
				d.ChopStream()
			}
			// debug
			//log.Println("debug:", d.ChoppedData)
			nextStep()
		}
	case 2: // "remap"
		{
			log.Println("state: ", d.ProcessID, state[d.State])
			errRemap := d.remapData()
			if errRemap != nil {
				d.errorList = append(d.errorList, errRemap)
				return errRemap
			}
			nextStep()
		}
	case 3: // "exportcsv"
		{
			log.Println("state: ", d.ProcessID, state[d.State])
			errCSV := d.ExportToCSV()
			if errCSV != nil {
				d.errorList = append(d.errorList, errCSV)
				return errCSV
			}
			nextStep()
		}
	case 4: // "exportxls"
		{
			log.Println("state: ", d.ProcessID, state[d.State])
			errXLS := d.ExportToXLS()
			if errXLS != nil {
				d.errorList = append(d.errorList, errXLS)
				return errXLS
			}
			nextStep()
		}
	case 5: // "combine"
		{
			log.Println("state: ", d.ProcessID, state[d.State])
			errCombine := d.Combiner()
			if errCombine != nil {
				log.Println("______________combine:", errCombine)
				d.errorList = append(d.errorList, errCombine)
				return errCombine
			}
			log.Println("combined:", d.combinedData)
			nextStep()
		}
	case 6: // "simplify"
		{
			log.Println("state: ", d.ProcessID, state[d.State])
			errSimplify := d.simplify()
			if errSimplify != nil {
				d.errorList = append(d.errorList, errSimplify)
				return errSimplify
			}
			log.Println("simplified:", d.combinedData)
			nextStep()
		}
	}
	return nil
}

// ChopCSV takes CSV file with data and extracts the selected positions as series.
func (d *Data) ChopCSV() error {
	readFile := func(pFilePath string) ([]string, error) {
		fileHandler, err := os.Open(pFilePath)
		if err != nil {
			return nil, err
		}
		defer fileHandler.Close()

		var result = make([]string, 0)
		scanner := bufio.NewScanner(fileHandler)
		for scanner.Scan() {
			result = append(result, scanner.Text())
		}
		return result, nil
	}

	content, errRead := readFile(d.InputCSV)
	if errRead != nil {
		return errRead
	}
	// pre - allocate result slices
	for r := range d.ChoppedData {
		d.ChoppedData[r] = make([]interface{}, len(content))
	}
	for k, v := range content {
		line := strings.Split(v, ",")
		if len(d.InputPositions) > (len(line)) {
			return errors.New(fmt.Sprintln("line", k, "has less positions"))
		}
		for _, p := range d.InputPositions {
			d.ChoppedData[p][k] = line[p]
		}
	}
	return nil
}

// ChopStream takes stream with data and extracts the selected positions as series.
func (d *Data) ChopStream() {
	// pre - allocate result slices
	log.Println("input stream:", d.InputStream, len(d.InputStream), d.InputPositions)
	for r := range d.ChoppedData {
		d.ChoppedData[r] = make([]interface{}, len(d.InputStream))
	}
	for k, v := range d.InputStream {
		//log.Println("input stream:", k, v)
		for k1, p := range d.InputPositions {
			//log.Println("pos:", p, d.ChoppedData[k1], v[p])
			d.ChoppedData[k1][k] = v[p]
		}
	}
	log.Println("chopped:", d.ChoppedData)
}

// remapData shifts data with 90 degrees.
func (d *Data) remapData() error {
	d.remappedData = make([][]interface{}, len(d.ChoppedData[0]))

	for i := 0; i < len(d.ChoppedData[0]); i++ {
		for j := 0; j < len(d.ChoppedData); j++ {
			buffer := d.ChoppedData[j][i]
			switch buffer.(type) {
			case string:
				{
					d.remappedData[i] = append(d.remappedData[i], buffer.(string))
				}
			case int:
				{
					b := strconv.Itoa(buffer.(int))
					d.remappedData[i] = append(d.remappedData[i], b)
				}
			case float64:
				{
					b := strconv.FormatFloat(buffer.(float64), 'f', 5, 64)
					d.remappedData[i] = append(d.remappedData[i], b)
				}
			}
		}
	}
	//log.Println("remapped:", d.remappedData)
	return nil
}

// ExportToCSV takes ChoppedData and spools it to CSV file set as ExportCSV.
func (d *Data) ExportToCSV() error {
	if2string := func(pIf []interface{}) []string {
		result := make([]string, len(pIf))
		for k, v := range pIf {
			switch v.(type) {
			case string:
				{
					result[k] = v.(string)
				}
			}
		}
		return result
	}

	csvFile, errCreate := os.Create(d.ExportCSV)
	defer csvFile.Close()
	if errCreate != nil {
		return errCreate
	}
	w := csv.NewWriter(csvFile)
	for _, record := range d.remappedData {
		errWrite := w.Write(if2string(record))
		if errWrite != nil {
			return errWrite
		}
	}
	w.Flush()

	errReported := w.Error()
	if errReported != nil {
		return errReported
	}
	return nil
}

// ExportToXLS takes ChoppedData and spools it to CSV file set as ExportCSV.
func (d *Data) ExportToXLS() error {
	x := exportxls.New([]string{}, d.remappedData, "result", d.ExportXLS)
	x.Auto()

	switch len(x.ErrorList) {
	case 0:
		{
			break
		}
	case 1:
		{
			return x.ErrorList[0]
		}
	default:
		{
			for range x.ErrorList {
				err := "Collected errors:\n"
				for i, e := range x.ErrorList {
					err = err + fmt.Sprintf("\tError %d: %s\n", i, e.Error())
				}
				return errors.New(err)
			}
		}
	}
	return nil
}

// Combiner takes remapped data, selects base and combines base with the other series.
func (d *Data) Combiner() error {
	if len(d.ChoppedData) < 2 {
		return errors.New("no series to combine with")
	}
	log.Println("base pos:", d.BasePosition, "chopped data:", d.ChoppedData)
	if len(d.ChoppedData[d.BasePosition]) == 0 {
		return errors.New("no values for base")
	}
	// transform to []float64 base series
	var baseSeries []float64
	for _, baseValue := range d.ChoppedData[d.BasePosition] {
		switch baseValue.(type) {
		case string:
			{
				walue, errParse := strconv.ParseFloat(baseValue.(string), 64)
				if errParse != nil {
					return errors.New("base series contains strings. aborting:" + baseValue.(string))
				}
				baseSeries = append(baseSeries, walue)
			}
		case int:
			{
				baseSeries = append(baseSeries, float64(baseValue.(int)))
			}
		case float64:
			{
				baseSeries = append(baseSeries, baseValue.(float64))
			}
		}

	}
	log.Println("baseSeries: ", baseSeries)
	// verify base series has the same number of entries as series to combine with
	for k := range d.ChoppedData {
		if len(baseSeries) != len(d.ChoppedData[k]) {
			return errors.New(fmt.Sprint("series ", k, " not balanced with base"))
		}
	}
	// combine now base with the remaining series from remapped data
	d.combinedData = make([][]simplify.Point, len(d.ChoppedData)-1)
	log.Println("combinedData:", d.combinedData)
	// pre - allocate result slices
	for r := range d.combinedData {
		d.combinedData[r] = make([]simplify.Point, len(baseSeries))
	}
	// merge slices
	var combineIndex int
	for k := range d.ChoppedData {
		// skip base series
		if k == d.BasePosition {
			continue
		}
		// fill series
		for i := range baseSeries {
			var buf simplify.Point
			buf.X = baseSeries[i]
			switch d.ChoppedData[k][i].(type) {
			case string:
				{
					walue, errParse := strconv.ParseFloat(d.ChoppedData[k][i].(string), 64)
					if errParse != nil {
						return errors.New("string series:" + d.ChoppedData[k][i].(string) + " aborting")
					}
					buf.Y = walue
				}
			case float64:
				{
					buf.Y = d.ChoppedData[k][i].(float64)
				}
			case int:
				{
					buf.Y = float64(d.ChoppedData[k][i].(int))
				}
			}
			d.combinedData[combineIndex][i] = buf
		}
		combineIndex++
	}
	return nil
}

func (d *Data) simplify() error {
	if len(d.SimplifyPositions) == 0 {
		return errors.New("no positions to simplify")
	}
	for _, s := range d.SimplifyPositions {
		if s >= len(d.remappedData) {
			return errors.New("simplify series outside data series")
		}
	}
	// now simplify series
	for _, r := range d.SimplifyPositions {
		b, errPoints := simplify.Points(d.combinedData[r], d.SimplifyTolerance, true)
		if errPoints != nil {
			return errPoints
		}
		d.combinedData[r] = b
	}
	return nil
}
