package main

import (
	"log"

	"github.com/PradeepPadmanabanC/golang_learning/learn_channels/helpers"
)

const numPool = 1000

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}

func main() {
	intChan := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)

	num := <-intChan
	log.Println(num)
}
