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

type ChartSerie struct {
	Name  string    `json:"name"`
	Type  string    `json:"type"` // spline, line, column
	YAxis int       `json:"yAxis"`
	Data  []float64 `json:"data"`
}

type Chart struct {
	Title               string
	Colors              []string
	YAxisTitle          string
	Width               int
	goTemplatePath      string // path to go template
	InterimJSONTemplate string // path to json config
	RenderSettings      []string
	ChartImage          string // path to png image file
	Series              []*ChartSerie
	PreparedSeries      string
	SupportedChartTypes map[string]struct{}
}

// NewChart - constructor
func NewChart(pTitle, pYAxisTitle, pGoTemplatePath, pInterimJSON, pRenderTo string, pWidth int) *Chart {
	instance := new(Chart)
	instance.Title = pTitle
	instance.YAxisTitle = pYAxisTitle
	instance.goTemplatePath = pGoTemplatePath
	instance.InterimJSONTemplate = pInterimJSON
	instance.ChartImage = pRenderTo

	// mapping chart types
	instance.SupportedChartTypes = make(map[string]struct{})
	instance.SupportedChartTypes["spline"] = struct{}{}
	instance.SupportedChartTypes["line"] = struct{}{}
	instance.SupportedChartTypes["column"] = struct{}{}

	instance.prepareRenderSettings()
	return instance
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

func (c *Chart) AddSerie(pSerie *ChartSerie) error {
	if len(pSerie.Name) == 0 {
		return errors.New("series error - no name")
	}
	if pSerie.YAxis != 0 {
		return errors.New("series error - invalid y axis")
	}
	if len(pSerie.Data) == 0 {
		return errors.New("series error - no data")
	}
	_, exists := c.SupportedChartTypes[pSerie.Type]
	if !exists {
		return errors.New("series error - invalid serie type")
	}
	c.Series = append(c.Series, pSerie)
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
	t, errParse := template.ParseFiles(c.goTemplatePath)
	if errParse != nil {
		log.Println("errParse:", errParse)
		return errParse
	}
	f, errCreate := os.Create(c.InterimJSONTemplate)
	defer f.Close()
	if errCreate != nil {
		log.Println("errCreate: ", errCreate)
		return errCreate
	}
	errExec := t.Execute(f, c)
	if errExec != nil {
		log.Println("errExec: ", errExec)
	}
	return errExec
}

// RenderChart - series need to be included
func (c *Chart) RenderChart() error {
	errTemplate := c.renderTemplate()
	if errTemplate != nil {
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
