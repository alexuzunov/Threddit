package handlers

import (
	"Threddit/internal/repositories"
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
		h.Home(w)
	})
	h.Get("/login", func(w http.ResponseWriter, r *http.Request) {
		users.LoginPage(w)
	})
	h.Get("/register", func(w http.ResponseWriter, r *http.Request) {
		users.RegisterPage(w)
	})
	h.Post("/api/register", func(w http.ResponseWriter, r *http.Request) {
		users.Register(w, r)
	})
	h.Post("/api/login", func(w http.ResponseWriter, r *http.Request) {
		users.Login(w, r)
	})

	h.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	return h
}

func (h *Handler) Home(w http.ResponseWriter) {
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
