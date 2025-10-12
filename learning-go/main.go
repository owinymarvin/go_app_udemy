package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJSON := `[
		{
		"first_name": "Clark",
		"last_name": "Kent", 
		"hair_color": "black",
		"has_dog": true
		}, 
		{
		"first_name": "Bruce",
		"last_name": "Wayne",
		"hair_color": "black",
		"has_dog": false
		}
	]`

	var unmarshalled []Person
	err := json.Unmarshal([]byte(myJSON), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling json", err)
	}
	log.Printf("unmarshalled: %v", unmarshalled)

	// write json from a struct
	var mySlice []Person
	var marshal1 Person
	marshal1.FirstName = "Diana"
	marshal1.LastName = "Prince"
	marshal1.HairColor = "red"
	marshal1.HasDog = false

	mySlice = append(mySlice, marshal1)

	var marshal2 Person
	marshal2.FirstName = "Kara"
	marshal2.LastName = "Zor-El"
	marshal2.HairColor = "blonde"
	marshal2.HasDog = true

	mySlice = append(mySlice, marshal2)

	newJSON, err := json.MarshalIndent(mySlice, "", "	")
	if err != nil {
		log.Println("Error marshalling json", err)
	}
	fmt.Println(string(newJSON) )
}
