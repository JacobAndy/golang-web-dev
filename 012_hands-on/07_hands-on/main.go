package main

import (
	"html/template"
	"log"
	"os"
)

type side struct {
	Name string
}

type nutrition struct {
	Calories int
	Carbs    int
	Fats     int
	Proteins int
}

type item struct {
	Name      string
	Price     int64
	Nutrition nutrition
	Sides     []side
}

type menu struct {
	Time  string
	Items []item
}

type restaurant struct {
	Name     string
	Location string
	Eat      []menu
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	data := restaurant{
		Name:     "Graces Kitchen",
		Location: "Downtown Nashville",
		Eat: []menu{
			menu{
				Time:  "Breakfast",
				Items: []item{item{"Bacon", 10, nutrition{100, 10, 10, 10}, []side{side{"Mac n Cheese"}, side{"Potatos"}, side{"Broccoli"}}}},
			},
		},
	}

	err := tpl.Execute(os.Stdout, data)

	if err != nil {
		log.Fatalln(err)
	}
}
