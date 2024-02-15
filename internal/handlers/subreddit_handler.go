package handlers

import (
	"Threddit/internal/helpers"
	"Threddit/internal/models"
	"Threddit/internal/repositories"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type SubredditHandler struct {
	repository *repositories.Repository
}

type SubredditJSON struct {
	Name string
	Type string
	NSFW string
}

func (h *SubredditHandler) SubredditCreatePage(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")

	if err != nil {
		http.Error(w, "Missing token", http.StatusUnauthorized)
		return
	}

	_, err = helpers.VerifyToken(cookie.Value)

	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	tmpl, err := template.ParseFiles(
		templateRoot+"sections.html",
		templateRoot+"subreddit_create.html",
		templateRoot+"navbar.html",
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "subreddit_create", nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *SubredditHandler) SubredditCreate(w http.ResponseWriter, r *http.Request) {
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

	var s SubredditJSON

	err = json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, _ := h.repository.FindByUsername(claims.Username)

	var subredditType models.SubredditType

	switch s.Type {
	case "public":
		subredditType = models.Public
	case "restricted":
		subredditType = models.Restricted
	case "private":
		subredditType = models.Private
	}

	nsfw, err := strconv.ParseBool(s.NSFW)
	if err != nil {
		log.Fatal(err)
	}

	subreddit := models.Subreddit{
		Name:      s.Name,
		Type:      subredditType,
		NSFW:      nsfw,
		CreatorID: user.ID,
	}

	if err := h.repository.CreateSubreddit(&subreddit); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
