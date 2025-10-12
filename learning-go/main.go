package main

import "log"

func main() {
	myMap := make(map[string]string)
	// first string is the index
	//second string is the value at that index
	myMap["dog"] = "Samson"
	log.Println("My Dogs name is ", myMap["dog"])
}
