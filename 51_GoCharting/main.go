package main

import (
	"bytes"
	"io/ioutil"
	"log"

	"github.com/wcharczuk/go-chart"
)

func main() {
	data := ChartData{
		xSeries:         []float64{1.0, 2.0, 3.0, 4.0},
		ySeries:         []float64{2.4, 5.6, 2.4, 5.6},
		xTitle:          "X Title",
		yTitle:          "Y Title",
		colorBackground: "928383",
		colorCanvas:     "bbadad",
		colorLine:       "302727",
		colorFill:       "ebe1e1",
	}

	c := NewChart(&data)
	buf := bytes.NewBuffer([]byte{})

	log.Println("error:", c.Render(chart.PNG, buf))
	log.Println("error:", ioutil.WriteFile("x.png", buf.Bytes(), 0644))

	/* 	e := email{
	   		from:    "sender@gmail.com",
	   		to:      []string{"to@gmail.com"},
	   		mime:    `MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n`,
	   		subject: `Subject: Test email from Go!\n`,
	   	}

	   	sendEmail(&e) */
}
