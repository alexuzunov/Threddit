package handlers

import (
	"html/template"
	"net/http"
)

func LoginPage(w http.ResponseWriter) {
	tmpl, err := template.ParseFiles("internal/templates/login.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = tmpl.Execute(w, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
