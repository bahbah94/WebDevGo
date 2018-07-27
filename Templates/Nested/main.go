package main

import (
	"html/template"
	"os"
	"log"
)

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main(){
	err := tpl.ExecuteTemplate(os.Stdout,"index.gohtml", 54)
	if err != nil{
		log.Fatalln(err)
	}
}
