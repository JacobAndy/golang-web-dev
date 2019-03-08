package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     int
	Region  string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	marriot := hotel{
		"Marriot",
		"Marriot lane",
		"Marriot",
		123,
		"Southern",
	}
	hilton := hotel{
		"Hilton",
		"Hilton lane",
		"Hilton",
		321,
		"Central",
	}
	data := []hotel{marriot, hilton}
	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
