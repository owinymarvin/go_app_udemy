package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/owinymarvin/go_app_udemy/pkg/config"
	"github.com/owinymarvin/go_app_udemy/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	return mux
}
