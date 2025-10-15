package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/owinymarvin/go_app_udemy/pkg/config"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

// make it start with an upercase letter to make it publicly available
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// get the template cache from the app config.
	tc := app.TemplateCache

	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from cache")
	}

	// declare a buffer to again thoroughly check through and see no erros are present
	//optional
	buf := new(bytes.Buffer)
	_ = t.Execute(buf, nil)
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	// myCache := make(map[string] *template.Template)
	myCache := map[string]*template.Template{}

	// get all the files labeled *.page.html from the /templates folder
	// pages,err := template.ParseGlob("./templates/*.page.html")
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	// range through all the files ending with *.page.html
	for _, page := range pages {
		name := filepath.Base(page)
		// ts short for template set
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			// adds the layout files to the templates folder.
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
