package main

import (
	"errors"
	"fmt"
	"net/http"
)

const PORTNUMBER string = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the Home Page")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	fmt.Fprintf(w, "This is the About Page \n")
	fmt.Fprintf(w, "sum is %d \n", sum)
}

func addValues(x, y int) int {
	return x + y
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 10.0)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")
		return
	}
	fmt.Fprintf(w, "%f divided by %f is %f", 100.0, 10.0, f)
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("Can't divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Printf("Application is running on port %s \n", PORTNUMBER)
	http.ListenAndServe(PORTNUMBER, nil)
}
