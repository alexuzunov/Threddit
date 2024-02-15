package handlers

import (
	"Threddit/internal/helpers"
	"Threddit/internal/models"
	"Threddit/internal/repositories"
	"encoding/json"
	"html/template"
	"net/http"
)

type PostHandler struct {
	repository *repositories.Repository
}

type PostJSON struct {
	Subreddit   string
	Title       string
	Description string
	Image       string
	ImageType   string
}

func (h *PostHandler) PostCreatePage(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")

	if err != nil {
		http.Error(w, "Missing token", http.StatusUnauthorized)
		return
	}

	claims, err := helpers.VerifyToken(cookie.Value)

	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	subreddits, _ := h.repository.GetCreatedSubreddits(claims.Username)

	tmpl, err := template.ParseFiles(
		templateRoot+"sections.html",
		templateRoot+"post_create.html",
		templateRoot+"navbar.html",
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "post_create", subreddits)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *PostHandler) PostCreate(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")

	if tokenString == "" {
		http.Error(w, "Missing authorization header", http.StatusUnauthorized)
		return
	}

	tokenString = tokenString[len("Bearer "):]

	claims, err := helpers.VerifyToken(tokenString)

	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	var p PostJSON

	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, _ := h.repository.FindByUsername(claims.Username)

	post := models.Post{
		Title:         p.Title,
		Text:          p.Description,
		AuthorID:      user.ID,
		SubredditName: p.Subreddit,
	}

	if err := h.repository.CreatePost(&post); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
