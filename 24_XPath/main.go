package main

import (
	"fmt"
	"strings"

	"github.com/antchfx/xquery/xml"
)

const url = "http://www.bnr.ro/nbrfxrates.xml"

// GetExchangeRate Returns exchange rate for passed currency symbol.
func GetExchangeRate(url, currency string) (string, error) {
	root, errLoad := xmlquery.LoadURL(url)
	if errLoad != nil {
		return "", errLoad
	}

	var rate string

	xmlquery.FindEach(root, "//Rate", func(i int, node *xmlquery.Node) {
		if strings.Contains(node.OutputXML(true), "\""+currency+"\"") {
			rate = node.FirstChild.InnerText()
		}
	})
	return rate, nil
}

func main() {
	currSymbol := "USD"
	rate, err := GetExchangeRate(url, currSymbol)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Curs:", "\""+currSymbol+"\"", rate)
}
