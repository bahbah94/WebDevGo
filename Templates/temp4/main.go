package main

import (
	"text/template"
	"os"
	"log"
)

var tpl *template.Template
func init(){
	tpl = template.Must(template.ParseFiles("myage.gohtml"))// so we're using must so we dont need to check for error
}
func main(){
	err := tpl.ExecuteTemplate(os.Stdout,"myage.gohtml", 23)
	if err != nil{
		log.Fatalln(err)
	}
}

