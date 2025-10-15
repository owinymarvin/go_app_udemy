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
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf("Application is running on port %s \n", PORTNUMBER)

	srv := &http.Server{
		Addr:    PORTNUMBER,
		Handler: routes(&app),
	}
	
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
