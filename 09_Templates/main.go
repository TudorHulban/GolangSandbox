package main

import (
	"log"
	"os"
	"text/template"
)

// Person Concentrates data for rendering the template.
type Person struct {
	Name  string
	Age   int
	Tasks []string
}

func main() {
	person := Person{
		Name:  "John",
		Age:   16,
		Tasks: []string{"T1", "T2", "T3"},
	}

	f1("template_name", person)
	f2("t2.tmpl", person)
	f3("t2.tmpl", "t2.csv", person)
}

// f1 Renders template to stdout.
func f1(templateName string, model Person) {
	t := template.New(templateName)

	t, errParse := t.Parse("hi {{.Name}}! rendering to stdout.")
	if errParse != nil {
		log.Println("errParse: ", errParse)
		return
	}
	t.Execute(os.Stdout, model)
}

// f2 Loads template and renders it to stdout.
func f2(templateFilePath string, model Person) {
	t, errParse := template.ParseFiles(templateFilePath)
	if errParse != nil {
		log.Println("errParse: ", errParse)
	}

	if errExec := t.Execute(os.Stdout, model); errExec != nil {
		log.Println("errExec: ", errExec)
	}
}

// f3 Loads template and renders it to file.
func f3(templateFilePath, renderTo string, model Person) {
	t, errParse := template.ParseFiles(templateFilePath)
	if errParse != nil {
		log.Println("errParse: ", errParse)
	}

	f, errCreate := os.Create(renderTo)
	if errCreate != nil {
		log.Println("errCreate: ", errCreate)
	}
	defer f.Close()

	if errExec := t.Execute(f, model); errExec != nil {
		log.Println("errExec: ", errExec)
	}
}
