package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastNAme  string
}

func main() {
	myMap := make(map[string]string)
	// first string is the index
	//second string is the value at that index
	myMap["dog"] = "Samson"
	log.Println("My Dogs name is ", myMap["dog"])

	myUserMap := make(map[string]User)

	me := User{
		FirstName: "Trevor",
		LastNAme:  "Mac Danniel",
	}

	myUserMap["me"] = me
	log.Println(myUserMap["me"].FirstName)
	log.Println(myUserMap["me"].LastNAme)

	// Maps are also immutable, hence very useful.

	var myNewVar float32
	myNewVar = 11.1
	log.Println(myNewVar)

	// if u dont know the datatype u will be storing in the map, make it open by using type interface ie
	// not recommended, but shown below.
	// myOpenMap := make(map[string]interface{})

	//SLICES, similar to arrays in other Programming languages

	var mySlice []string
	mySlice = append(mySlice, "Dmarv")
	mySlice = append(mySlice, "Timothy")
	mySlice = append(mySlice, "Owiny")

	log.Println("My slice has ", mySlice)
	sort.Strings(mySlice)
	log.Println("My slice has been sorted to", mySlice)

	// more slices

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	log.Println("The numbers in slice are: ", numbers)

	log.Println("Printing the first 3 numbers in a slice: ", numbers[0:3])

}
