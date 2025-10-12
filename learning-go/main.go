package main

import "fmt"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		Name:  "Samson",
		Breed: "German Shepherd",
	}
	gorilla := Gorilla{
		Name:          "Jock",
		Color:         "Black",
		NumberOfTeeth: 32,
	}

	PrintInfo(&dog)
	PrintInfo(&gorilla)

}

func PrintInfo(a Animal) {
	fmt.Println("This animal says ", a.Says(), " and has ", a.NumberOfLegs(), " legs")
}

func (d *Dog) Says() string {
	return "Woof! woof!"
}
func (d *Dog) NumberOfLegs() int {
	return 4
}
func (d *Gorilla) Says() string {
	return "King Kong"
}
func (d *Gorilla) NumberOfLegs() int {
	return 8
}
