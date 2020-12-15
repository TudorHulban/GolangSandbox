package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"log"
	"net/http"
	"time"

	datagen "github.com/TudorHulban/DataGenerator"
	simplex "github.com/TudorHulban/LineSimplification"
)

type ChartTitle struct {
	Display bool   `json:"display"`
	Text    string `json:"text"`
}

type ChartOptions struct {
	Responsive bool       `json:"responsive"`
	Title      ChartTitle `json:title`
}

type ChartDataset struct {
	Data        []float64 `json:"data"`
	Label       string    `json:"label"`
	Fill        bool      `json:"fill"`
	BorderColor string    `json:"borderColor"`
}

type ChartData struct {
	Labels   []float64      `json:"labels"`
	Datasets []ChartDataset `json:"datasets"`
}

type ChartConfig struct {
	Type    string       `json:"type"`
	Data    ChartData    `json:"data"`
	Options ChartOptions `json:"options"`
}

//var chartData ChartResponse

// chartLabels - labels common to raw and optimized points
var chartLabels []float64

// simplexData - data points after simplification
var simplexData []float64

const port = 3000
const jsonConfig = `{"config":[{"name":"a","ftype":2, "flength":5, "fpositive":true, "min":0, "max":100},{"name":"b","ftype":0, "flength":5, "fpositive":true, "min":0, "max":100}], "norows":200}`

func main() {
	http.HandleFunc("/raw", chartRaw)
	http.HandleFunc("/simple", chartSimple)
	http.HandleFunc("/", indexHandler)

	http.ListenAndServe((fmt.Sprintf(":%v", port)), nil)
}

func chartRaw(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	dataConfig, _ := datagen.GetConfig(jsonConfig)
	dataChart, _ := datagen.GetData(dataConfig)
	dataFloat, _ := SliceInterfToFloat(&dataChart.Rows)
	log.Println("chartRaw | data:", dataFloat)

	simpleData, _ := simplex.Simplify(*dataFloat, 3, false)
	_, simplexData, _ = SplitPointsSlice(simpleData)
	log.Println("raw:", dataFloat)

	data := []float64{}
	chartLabels, data, _ = SplitPointsSlice(*dataFloat)
	response := NewChart(chartLabels, data)

	json2stream := json.NewEncoder(w)
	json2stream.Encode(&response)

	log.Println("y:", time.Now().Sub(startTime))
}

func chartSimple(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	response := NewChart(chartLabels, simplexData)
	json2stream := json.NewEncoder(w)
	json2stream.Encode(&response)

	log.Println("y:", time.Now().Sub(startTime))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}

func NewChart(labels, data []float64) *ChartConfig {
	// https://tobiasahlin.com/blog/chartjs-charts-to-get-you-started/

	return &ChartConfig{
		Type: "line",
		Options: ChartOptions{
			Responsive: true,
			Title: ChartTitle{
				Display: true,
				Text:    "Chart 1",
			},
		},
		Data: ChartData{
			Labels: labels,
			Datasets: []ChartDataset{ChartDataset{
				Label:       "xxx",
				Fill:        false,
				BorderColor: "#3e95cd",
				Data:        data,
			}},
		},
	}
}

// SplitPointsSlice - returns sliceLabels, sliceData
func SplitPointsSlice(slice [][]float64) ([]float64, []float64, error) {
	// assumption - input is slice of Points, aka 2 values
	sliceLabels := []float64{}
	sliceData := []float64{}

	for k, v := range slice {
		if len(v) != 2 {
			return nil, nil, errors.New("Slices has issues at row: " + strconv.FormatInt(int64(k), 10))
		}
		sliceLabels = append(sliceLabels, v[0])
		sliceData = append(sliceData, v[1])
	}
	return sliceLabels, sliceData, nil
}

func SliceInterfToFloat(slice *[][]interface{}) (*[][]float64, error) {
	instance := make([][]float64, 0)

	for _, v := range *slice {
		row := []float64{}

		for _, vv := range v {
			switch vv.(type) {
			case float64:
				row = append(row, vv.(float64))
			case int:
				row = append(row, float64(vv.(int)))
			}
		}
		instance = append(instance, row)
	}
	return &instance, nil
}
