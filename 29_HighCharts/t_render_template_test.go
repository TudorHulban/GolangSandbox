package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const chartType = "spline"

func TestRenderGoTemplate(t *testing.T) {
	c := NewChart("Title of Chart", "Y Title", "chartConfig.tmpl", "go_rendered.json", "x.png", 500)

	assert.Nil(t, c.renderTemplate())

	s1 := ChartSerie{
		Name:  "S1",
		Type:  chartType,
		YAxis: 0,
		Data:  []float64{1.0, 2.4, 4.5},
	}

	s2 := ChartSerie{
		Name:  "S2",
		Type:  chartType,
		YAxis: 0,
		Data:  []float64{3.0, 5.4, 7.5},
	}

	s3 := ChartSerie{
		Name:  "S3",
		Type:  chartType,
		YAxis: 0,
		Data:  []float64{6.0, 8.4, 14.5},
	}

	c.AddSerie(&s1)
	c.AddSerie(&s2)
	c.AddSerie(&s3)

	c.prepareSeries()

	assert.Nil(t, c.RenderChart())
}
