package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"os"
	"os/exec"
	"strconv"
	"syscall"
	"text/template"
)

const nodePath = "/home/tudi/.config/versions/node/v12.6.0/bin/highcharts-export-server"

// ChartSerie - Structure used for variables to hold chart series.
type ChartSerie struct {
	Name  string    `json:"name"`
	Type  string    `json:"type"` // spline, line, column
	YAxis int       `json:"yAxis"`
	Data  []float64 `json:"data"`
}

type Chart struct {
	Title               string
	YAxisTitle          string
	Width               int
	GoTemplatePath      string // path to go template
	InterimJSONTemplate string // path to json config
	ChartImage          string // path to png image file
	PreparedSeries      string
	Colors              []string
	RenderSettings      []string
	Series              []*ChartSerie
	SupportedChartTypes map[string]struct{}
}

// NewChart - constructor
// TODO: add create config helper.
func NewChart(title, yAxisTitle, goTemplatePath, interimJSON, renderTo string, width int) *Chart {
	instance := Chart{
		Title:               title,
		YAxisTitle:          yAxisTitle,
		GoTemplatePath:      goTemplatePath,
		InterimJSONTemplate: interimJSON,
		ChartImage:          renderTo,
	}

	// mapping chart types
	instance.SupportedChartTypes = make(map[string]struct{})
	instance.SupportedChartTypes["spline"] = struct{}{}
	instance.SupportedChartTypes["line"] = struct{}{}
	instance.SupportedChartTypes["column"] = struct{}{}

	instance.prepareRenderSettings()
	return &instance
}

// prepareRenderSettings - private
func (c *Chart) prepareRenderSettings() {
	c.RenderSettings = append(c.RenderSettings, "") // needed, looks like a bug with highcharts
	c.RenderSettings = append(c.RenderSettings, "-infile")
	c.RenderSettings = append(c.RenderSettings, c.InterimJSONTemplate)
	c.RenderSettings = append(c.RenderSettings, "-outfile")
	c.RenderSettings = append(c.RenderSettings, c.ChartImage)
	c.RenderSettings = append(c.RenderSettings, "-width")
	c.RenderSettings = append(c.RenderSettings, strconv.Itoa(c.Width))
}

func (c *Chart) AddSerie(serie *ChartSerie) error {
	if len(serie.Name) == 0 {
		return errors.New("series error - no name")
	}
	if serie.YAxis != 0 {
		return errors.New("series error - invalid y axis")
	}
	if len(serie.Data) == 0 {
		return errors.New("series error - no data")
	}

	if _, exists := c.SupportedChartTypes[serie.Type]; !exists {
		return errors.New("series error - invalid serie type")
	}

	c.Series = append(c.Series, serie)
	return nil
}

// prepareSeries - index is for series position in chart slice. marshals series structure to json
func (c *Chart) prepareSeries() {
	for _, v := range c.Series {
		var b bytes.Buffer

		j := json.NewEncoder(&b)
		j.Encode(&v)
		c.PreparedSeries = c.PreparedSeries + "," + b.String()
	}

	c.PreparedSeries = c.PreparedSeries[1:]
	log.Println("prepared series:", c.PreparedSeries)
}

func (c *Chart) renderTemplate() error {
	if len(c.Series) == 0 {
		return errors.New("no series")
	}

	t, errParse := template.ParseFiles(c.GoTemplatePath)
	if errParse != nil {
		log.Println("errParse:", errParse)
		return errParse
	}

	f, errCreate := os.Create(c.InterimJSONTemplate)
	if errCreate != nil {
		log.Println("errCreate: ", errCreate)
		return errCreate
	}
	defer f.Close()

	if errExec := t.Execute(f, c); errExec != nil {
		log.Println("errExec: ", errExec)
		return errExec
	}

	return nil
}

// RenderChart - series need to be included
func (c *Chart) RenderChart() error {
	if errTemplate := c.renderTemplate(); errTemplate != nil {
		return errTemplate
	}

	binary, errPath := exec.LookPath(nodePath)
	if errPath != nil {
		log.Println("look path error", errPath)
		return errPath
	}

	errExec := syscall.Exec(binary, c.RenderSettings, os.Environ())
	if errExec != nil {
		log.Println("exec error", errExec)
		return errExec
	}

	return nil
}
