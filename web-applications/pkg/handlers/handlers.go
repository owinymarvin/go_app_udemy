package handlers

import (
	"log"
	"net/http"

	"github.com/owinymarvin/go_app_udemy/pkg/config"
	"github.com/owinymarvin/go_app_udemy/pkg/models"
	"github.com/owinymarvin/go_app_udemy/pkg/render"
)

// creating a repository pattern
type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
	log.Println("Request URL:", r.URL.Path)
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	log.Println("Request URL:", r.URL.Path)
	// perform some logic
	stringMap := map[string]string{}
	stringMap["test"] = "Hello, again."

	// send data to the template
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{StringMap: stringMap})
}
