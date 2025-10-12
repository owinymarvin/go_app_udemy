package main

import "log"

func main() {
	animals := []string{"dog", "cat", "bird", "fish", "horse"}
	for i, animal := range animals {
		log.Println(i, animal)
	}

	mapedAnimals := make(map[string]string)
	mapedAnimals["dog"] = "Fideo"
	mapedAnimals["cat"] = "murphy"

	for mapedAnimalType, mapedAnimal := range mapedAnimals {
		log.Println(mapedAnimalType, " => ", mapedAnimal)
	}

	// ranging over a string
	var firstLine = "Once upon a midnight dreary"
	log.Println(firstLine)
	for index, letter := range firstLine {
		log.Println(index, ":", letter)
	}

	//ranging over a struct
	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
		Password  string
	}
	// var users []User
	var users []User
	users = append(users, User{"John", "Smith", "XXXXXXXXXXXXXX", 30, "password"})
	users = append(users, User{"Mary", "Smith", "YYYYYYYYYYYYYY", 25, "password2"})
	users = append(users, User{"Sally", "Smith", "ZZZZZZZZZZZZZZ", 40, "password3"})
	users = append(users, User{"Alex", "Smith", "AAAAAAAAAAAAAAA", 45, "password4"})
	users = append(users, User{"Michael", "Smith", "BBBBBBBBBBBBBBB", 35, "password5"})

	for i, l := range users {
		log.Println(i, "=>", l.FirstName, l.LastName, l.Email, l.Age, l.Password)
	}
}
