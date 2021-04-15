package main

import (
	"embed"
	"fmt"
	"html/template"
	"os"
)

type User struct {
	IsAdmin bool
	Name    string
}

//go:embed templates/*.gohtml
var f embed.FS

func main() {
	funcMap := template.FuncMap{
		"isAdmin": func(isAdmin bool) bool {
			return isAdmin
		},
	}

	tmpl := template.New("views")
	tmpl.Funcs(funcMap)

	tmpl, err := tmpl.ParseFS(f, "templates/*.gohtml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	tmpl.ExecuteTemplate(os.Stdout, "r_layout.gohtml", User{
		IsAdmin: false,
		Name:    "john"})
}
