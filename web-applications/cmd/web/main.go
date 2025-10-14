package main

import (
	"fmt"
	"net/http"
	"github.com/owinymarvin/go_app_udemy/pkg/handlers"
)

const PORTNUMBER string = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Application is running on port %s \n", PORTNUMBER)
	err := http.ListenAndServe(PORTNUMBER, nil)
	if err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
