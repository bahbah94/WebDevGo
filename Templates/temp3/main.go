package main

import (
	"text/template"
	"log"
	"os"
)


func main(){
	tpl, err := template.ParseGlob("templates/*.gohtml")
	if err != nil{
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout,nil)
	if err != nil{
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil{
		log.Fatalln(err)
	}
}
