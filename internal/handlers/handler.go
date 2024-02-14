package handlers

import (
	"Threddit/internal/repositories"
	"errors"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"html/template"
	"net/http"
)

type Handler struct {
	*chi.Mux

	repository *repositories.Repository
}

const templateRoot = "internal/templates/"

func NewHandler(repository *repositories.Repository) *Handler {
	h := &Handler{
		Mux:        chi.NewMux(),
		repository: repository,
	}

	users := UserHandler{repository: repository}

	h.Use(middleware.Logger)

	h.Get("/", func(w http.ResponseWriter, r *http.Request) {
		h.Home(w, r)
	})
	h.Get("/login", func(w http.ResponseWriter, r *http.Request) {
		users.LoginPage(w, r)
	})
	h.Get("/register", func(w http.ResponseWriter, r *http.Request) {
		users.RegisterPage(w, r)
	})
	h.Post("/api/users", func(w http.ResponseWriter, r *http.Request) {
		users.Register(w, r)
	})
	h.Post("/api/login", func(w http.ResponseWriter, r *http.Request) {
		users.Login(w, r)
	})

	h.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	return h
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	logged := true
	_, err := r.Cookie("token")

	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			logged = false
		}
	}

	data := map[string]interface{}{
		"Logged": logged,
	}

	tmpl, err := template.ParseFiles(
		templateRoot+"sections.html",
		templateRoot+"home.html",
		templateRoot+"navbar.html",
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = tmpl.ExecuteTemplate(w, "home", data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
