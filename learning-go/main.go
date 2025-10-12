package main

import (
	"log"

	"github.com/owinymarvin/go_app_udemy/helpers"
)

const constNumPool = 10

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(constNumPool)
	intChan <- randomNumber

}

func main() {
	// channel, received in 1 or more places in program
	intChannel := make(chan int)
	defer close(intChannel) // closes the channel to prevent memory leaks

	// concurrent operation, or go routine.
	// A go routines run at the same time. So concurrency
	//calls the function defined above
	go CalculateValue(intChannel)

	num := <-intChannel // waits for the channel to return a value
	log.Println(num)
}
