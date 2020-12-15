package main

import (
	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
)

// ChartData Type that consolidates chart data.
type ChartData struct {
	xSeries []float64
	ySeries []float64
	xTitle  string
	yTitle  string

	//hex, css like
	colorBackground string
	colorCanvas     string
	colorLine       string
	colorFill       string
}

// NewChart Constructor for chart.
func NewChart(chartData *ChartData) chart.Chart {
	return chart.Chart{
		Background: chart.Style{
			FillColor: drawing.ColorFromHex(chartData.colorBackground),
		},
		Canvas: chart.Style{
			FillColor: drawing.ColorFromHex(chartData.colorCanvas),
		},
		XAxis: chart.XAxis{
			Name: chartData.xTitle,
			GridMinorStyle: chart.Style{
				StrokeColor: drawing.Color{R: 0, G: 0, B: 0, A: 100},
				StrokeWidth: 1.0,
			},
		},
		YAxis: chart.YAxis{
			Name: chartData.yTitle,
			GridMinorStyle: chart.Style{
				StrokeColor: drawing.Color{R: 0, G: 0, B: 0, A: 100},
				StrokeWidth: 1.0,
			},
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Style: chart.Style{
					StrokeColor: drawing.ColorFromHex(chartData.colorLine),
					StrokeWidth: 9,
					FillColor:   drawing.ColorFromHex(chartData.colorFill),
				},
				XValues: chartData.xSeries,
				YValues: chartData.ySeries,
			}},
	}
}
