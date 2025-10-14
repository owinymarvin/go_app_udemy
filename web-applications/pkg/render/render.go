package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// make it start with an upercase letter to make it publicly available
func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
	parsedTemplate, err := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println("Error parsing the template:", err)
		return
	}

	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println("Error executing the template:", err)
		return
	}
}

// tc = template cache
var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// check to see if we already have the template in the cache
	tmpl, inMap := tc[t]
	if !inMap {
		// need to create the template
		log.Println("creating template and adding to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		// get the template from cache after creation
		tmpl = tc[t]
	} else if inMap {
		// we already have the template in the cache
		log.Println("using cached template")
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.html",
	}
	// parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	// add template to cache
	tc[t] = tmpl
	return nil
}
