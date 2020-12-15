package main

import (
	"testing"
)

func TestRenderGoTemplate(t *testing.T) {
	c := NewChart("Title of Chart", "Y Title", "chartConfig.tmpl", "go_rendered.json", "x.png", 500)
	errRenderTempl := c.renderTemplate()
	if errRenderTempl != nil {
		t.Error(errRenderTempl)
	}
	var s1 ChartSerie
	s1.Name = "S1"
	s1.Type = "spline1"
	s1.YAxis = 0
	s1.Data = []float64{1.0, 2.4, 4.5}

	var s2 ChartSerie
	s2.Name = "S2"
	s2.Type = "spline"
	s2.YAxis = 0
	s2.Data = []float64{3.0, 5.4, 7.5}

	var s3 ChartSerie
	s3.Name = "S3"
	s3.Type = "spline"
	s3.YAxis = 0
	s3.Data = []float64{6.0, 8.4, 14.5}

	c.AddSerie(&s1)
	c.AddSerie(&s2)
	c.AddSerie(&s3)
	c.prepareSeries()

	errRenderChart := c.RenderChart()
	if errRenderChart != nil {
		t.Error(errRenderChart)
	}
}
