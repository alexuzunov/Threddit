package handlers

import (
	"html/template"
	"net/http"
)

const templateRoot = "internal/templates/"

func Home(w http.ResponseWriter) {
	tmpl, err := template.ParseFiles(
		templateRoot+"sections.html",
		templateRoot+"home.html",
		templateRoot+"navbar.html",
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = tmpl.ExecuteTemplate(w, "home", nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
