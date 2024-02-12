package main

import (
	"Threddit/internal/database"
	"Threddit/internal/handlers"
	"Threddit/internal/models"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
)

func main() {
	db, err := database.ConnectToDatabase()

	if err != nil {
		log.Fatal(fmt.Sprintf("Error opening connection: %s", err.Error()))
	}

	err = db.AutoMigrate(&models.User{}, &models.Subreddit{}, &models.Post{}, &models.Comment{}, &models.Vote{})
	if err != nil {
		log.Fatal(fmt.Sprintf("Error during migration: %s", err.Error()))
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
		handlers.LoginPage(w)
	})

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error during server initiation: %s", err.Error()))
	}
}
