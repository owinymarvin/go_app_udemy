package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main() {
	var myVar myStruct
	myVar.FirstName = "John"
	myVar2 := myStruct{
		FirstName: "Mary",
	}
	log.Println("my Var is ", myVar.printFirstName())
	log.Println("my Var2 is ", myVar2.printFirstName())
}
