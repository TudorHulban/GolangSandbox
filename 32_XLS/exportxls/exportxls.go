package exportxls

import (
	"log"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/pkg/errors"
)

var alphabet = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

// Cfg Holds configuration for constructor.
type Cfg struct {
	StyleColumnNames int
	StyleCellsData   int
	SheetName        string
	SaveTo           string   // path for spool
	ColumnNames      []string // if nil sheet will not have column names.
}

// XLS Provides support for configuration and logic.
type XLS struct {
	Cfg
	step       int // internal step in flow
	firstRow   int
	RowsData   [][]interface{}
	handlerXLS *excelize.File
	ErrorList  []error // accumulated errors
}

// steps Drafts the steps in the flow.
func steps(ix int) string {
	step := []string{
		"Init",
		"ValidateData",
		"SetStyle",
		"SetColumnsData",
		"SetRowsData",
		"Save",
		"Stop",
	}

	if (ix > len(step)-1) || (ix < 0) {
		return ""
	}

	return step[ix]
}

// NewXLS is constructor style.
func NewXLS(cfg Cfg, data [][]interface{}) *XLS {
	return &XLS{
		Cfg:      cfg,
		step:     0,
		firstRow: 2,
		RowsData: data,
	}
}

// Process drives the flow forward. Provides clear picture of the flow steps.
// Each step with logging.
// Errors are added to slice as due to reccurence the error would not be directly returned.
func (x *XLS) Process() error {
	nextStep := func() {
		x.step++
		x.Process()
	}

	switch x.step {
	case 0: // "init"
		{
			log.Println("case 0 - step: ", steps(x.step))

			if errInit := x.init(); errInit != nil {
				x.ErrorList = append(x.ErrorList, errInit)
				return errInit
			}
			nextStep()
		}

	case 1: // "validate_data"
		{
			log.Println("case 1 - step: ", steps(x.step))

			if errValidate := x.validateData(); errValidate != nil {
				x.ErrorList = append(x.ErrorList, errValidate)
				log.Println("validation failed: ", errValidate)
				return errValidate
			}

			nextStep()
		}

	case 2: // "setstyle"
		{
			if len(x.ColumnNames) == 0 {
				log.Println("case 2 - step: ", "ColumnNames is nil")
				nextStep()
				return nil
			}

			log.Println("case 2 - step: ", steps(x.step))

			if errStyle := x.setStyle(); errStyle != nil {
				x.ErrorList = append(x.ErrorList, errStyle)
				return errStyle
			}

			nextStep()
		}

	case 3: // "setcolumnsdata"
		{
			if len(x.ColumnNames) == 0 {
				log.Println("case 3 - step: ", "ColumnNames is nil")
				nextStep()
				return nil
			}

			log.Println("case 3 - step: ", steps(x.step))

			if errColumns := x.setColumnData(); errColumns != nil {
				x.ErrorList = append(x.ErrorList, errColumns)
				return errColumns
			}

			nextStep()
		}

	case 4: // "setrowsdata"
		{
			log.Println("case 4 - step: ", steps(x.step))

			if errRows := x.setRowsData(); errRows != nil {
				x.ErrorList = append(x.ErrorList, errRows)
				return errRows
			}
			nextStep()
		}

	case 5: // "save"
		{
			log.Println("case 5 - step: ", steps(x.step))

			if errSave := x.exportToFile(); errSave != nil {
				x.ErrorList = append(x.ErrorList, errSave)
				return errSave
			}
			nextStep()
		}

	default: // "stop" {
		log.Println("case default - step: ", steps(x.step))
	}
	return nil
}

// Init Method initializing the xls handler, sheet and styles.
func (x *XLS) init() error {
	x.handlerXLS = excelize.NewFile()
	x.handlerXLS.NewSheet(x.SheetName)
	x.handlerXLS.DeleteSheet("Sheet1")

	styleColumns, errStyle := x.handlerXLS.NewStyle(`{"font":{"bold":true}, "alignment":{"horizontal":"center"}, "fill":{"type":"pattern", "color":["#e0ebf5"], "pattern":1}}`)
	if errStyle != nil {
		return errors.WithMessagef(errStyle, "step: %v", steps(x.step))
	}
	x.StyleColumnNames = styleColumns

	styleCells, errStyle := x.handlerXLS.NewStyle(`{"alignment":{"horizontal":"center"}}`)
	if errStyle != nil {
		return errors.WithMessagef(errStyle, "step: %v", steps(x.step))
	}

	x.StyleCellsData = styleCells
	return nil
}

// validateData validates if xls data passes validation criterias.
func (x *XLS) validateData() error {
	if len(x.ColumnNames) > 26 {
		return errors.WithMessagef(errors.New("column names exceeding maximum value"), "step: %v", steps(x.step))
	}

	if len(x.RowsData) == 0 {
		return errors.WithMessagef(errors.New("no row data"), "step: %v", steps(x.step))
	}

	if len(x.ColumnNames) == 0 {
		x.firstRow = 1 // no header
		return nil
	}

	if len(x.ColumnNames) != len(x.RowsData[0]) {
		//return errors.WithMessagef(errors.New("column names not in sync with row data"), "step: %v", steps(x.step))
		return errors.New("column names not in sync with row data")
	}

	return nil
}

func (x *XLS) setStyle() error {
	x.handlerXLS.SetCellStyle(x.SheetName, alphabet[0]+"1", alphabet[len(x.ColumnNames)-1]+"1", x.StyleColumnNames)

	// set columns bold
	x.handlerXLS.SetCellStyle(x.SheetName, "A2", alphabet[len(x.ColumnNames)-1]+strconv.Itoa(len(x.RowsData[0])+1), x.StyleCellsData)

	return nil
}

func (x *XLS) setColumnData() error {
	for k, v := range x.ColumnNames {
		x.handlerXLS.SetCellValue(x.SheetName, alphabet[k]+"1", v)
	}

	return nil
}

func (x *XLS) setRowsData() error {
	log.Println("Data:", x.RowsData)

	for ixRow, row := range x.RowsData {
		for ixCell, valCell := range row {
			x.handlerXLS.SetCellValue(x.SheetName, alphabet[ixCell]+strconv.Itoa(ixRow+x.firstRow), valCell)
		}
	}

	return nil
}

func (x *XLS) exportToFile() error {
	if errSave := x.handlerXLS.SaveAs(x.SaveTo); errSave != nil {
		return errors.WithMessagef(errSave, "step: %v", steps(x.step))
	}

	return nil
}
