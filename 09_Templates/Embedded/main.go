package main

import (
	"embed"
	"fmt"
	"html/template"
	"os"
)

type Session struct {
	IsAuthenticated bool
	Username        string
}

type State struct {
	Session
	Data []string
}

const chosenFruit = "Banana"

//go:embed templates/*.gohtml
var f embed.FS

func main() {
	funcMap := template.FuncMap{
		"isChosen": func(fruit string) bool {
			if fruit == chosenFruit {
				return true
			}

			return false
		},
	}

	tmpl := template.New("")
	tmpl.Funcs(funcMap)

	tmpl, err := tmpl.ParseFS(f, "templates/index.gohtml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	session := Session{
		IsAuthenticated: true,
		Username:        "john",
	}
	data := []string{"Orange", "Apple", "Banana"}

	tmpl.ExecuteTemplate(os.Stdout, "index.gohtml", State{
		Session: session,
		Data:    data,
	})
}
