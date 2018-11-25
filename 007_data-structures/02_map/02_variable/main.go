package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("tpl.gohtml"))
}

func main() {
	m := map[string]string{
		"India":    "Gandhi",
		"America":  "MLK",
		"Meditate": "Buddha",
		"Love":     "Jesus",
		"Prophet":  "Mohammad",
	}

	err := tpl.Execute(os.Stdout, m)
	if err != nil {
		log.Fatal(err)
	}
}
