package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// make it start with an upercase letter to make it publicly available
func RenderTemplate(w http.ResponseWriter, tmpl string) {
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
