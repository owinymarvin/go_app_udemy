package main

import (
	"log"

	"github.com/owinymarvin/go_app_udemy/helpers"
)

func main() {
	var myVar helpers.SomeType
	myVar.TypeName = "Some name"
	myVar.TypeNumber = 10

	log.Println(myVar.TypeName)
	log.Println(myVar.TypeNumber)
}
