package main

import (
	"html/template"
	"os"
	"log"
)

var tpl *template.Template

type sage struct{
	Name string
	Motto string
}
func init(){
	tpl = template.Must(template.ParseFiles(("stuff.gohtml")))
}
func main(){
	buddha := sage{
		"Buddha",
		"The belief of shit",
	}
	err := tpl.Execute(os.Stdout, buddha)
	if err!= nil{
		log.Fatalln(err)
	}
}