package main

import (
	"log"

	"github.com/PradeepPadmanabanC/golang_learning/learn_packages/helpers"
)

func main() {
	log.Println("Hello")
	var myVar helpers.SomeType
	myVar.TypeName = "Some Name"
	log.Println(myVar.TypeName)
}
