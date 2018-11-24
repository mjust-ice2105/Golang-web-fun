package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("index.html")
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.Execute(file, nil)
	if err != nil {
		log.Fatal(err)
	}

}
