package main

import (
	"errors"
	"log"
)

func main() {
	result, err := divide(100, 10)
	if err != nil {
		log.Println(err.Error())
		log.Println(err)
		return
	}
	log.Println("result of my division", result)

}

func divide(x, y float32) (float32, error) {
	var result float32

	if y == 0 {
		return 0, errors.New("can't divide by zero")
	}

	result = x / y
	return result, nil
}
