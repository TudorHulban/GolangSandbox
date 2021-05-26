package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func scrape() {
	res, err := http.Get("https://www.w3schools.com/w3css/tryw3css_examples_login.htm")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("label").Each(func(i int, s *goquery.Selection) {
		fmt.Println(i, s.Text())
	})

	doc.Find("button").Each(func(i int, s *goquery.Selection) {
		fmt.Println(i, s.Text())
	})
	
	nodes := doc.Find("div .top10 a").Nodes

	for _, node := range nodes {
		links := node.Attr

		for _, link := range links {
			if strings.Contains(link.Val, "sql") {
				fmt.Println(link)
				break
			}
		}
	}
}

func main() {
	scrape()
}
