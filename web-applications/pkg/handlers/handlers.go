package handlers

import (
	"log"
	"net/http"

	"github.com/owinymarvin/go_app_udemy/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")
	log.Println("Request URL:", r.URL.Path)
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")
	log.Println("Request URL:", r.URL.Path)

}
