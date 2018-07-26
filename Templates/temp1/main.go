package main

import "fmt"

func main(){
	name := "Randall"

	tpl := `
		<!DOCTYPE html>
			<html lang ="en">
			<head>
			<meta charset ="UTF-8">
				<title>Hello words!</title>
			</head>
			<body>
				<h1>` + name + `</h1>
				</body>
				</html>`

				fmt.Println(tpl)
}
