package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/owinymarvin/go_app_udemy/pkg/config"
	"github.com/owinymarvin/go_app_udemy/pkg/handlers"
	"github.com/owinymarvin/go_app_udemy/pkg/render"
)

const PORTNUMBER string = ":8080"

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Application is running on port %s \n", PORTNUMBER)
	http.ListenAndServe(PORTNUMBER, nil)
}
