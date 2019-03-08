package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

type eat struct {
	Time    string
	Serving []string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	b := eat{
		Time:    "Breakfast",
		Serving: []string{"Eggs", "Bacon", "Sausage"},
	}
	l := eat{
		Time:    "Lunch",
		Serving: []string{"Roast Beef Sandwich", "Burrito", "Pizza"},
	}
	d := eat{
		Time:    "Dinner",
		Serving: []string{"Tomahawk Steak", "Filet", "Salad Bar"},
	}
	r := []eat{b, l, d}
	err := tpl.Execute(os.Stdout, r)
	if err != nil {
		log.Fatalln(err)
	}
}
