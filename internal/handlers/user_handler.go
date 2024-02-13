package handlers

import (
	"encoding/json"
	"html/template"
	"net/http"
)

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterForm struct {
	Email           string `json:"email"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func LoginPage(w http.ResponseWriter) {
	tmpl, err := template.ParseFiles(
		templateRoot+"sections.html",
		templateRoot+"login.html",
		templateRoot+"navbar.html",
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = tmpl.ExecuteTemplate(w, "login", nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	var form LoginForm

	err := json.NewDecoder(r.Body).Decode(&form)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

}

func RegisterPage(w http.ResponseWriter) {
	tmpl, err := template.ParseFiles(
		templateRoot+"sections.html",
		templateRoot+"register.html",
		templateRoot+"navbar.html",
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = tmpl.ExecuteTemplate(w, "register", nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	var form RegisterForm

	err := json.NewDecoder(r.Body).Decode(&form)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
