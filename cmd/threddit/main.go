package main

import (
	"Threddit/internal/database"
	"Threddit/internal/handlers"
	"Threddit/internal/models"
	"Threddit/internal/repositories"
	"fmt"
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

	repository, err := repositories.NewRepository(db)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error during repository initialization: %s", err.Error()))
	}

	h := handlers.NewHandler(repository)

	err = http.ListenAndServe(":8080", h)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error during server initiation: %s", err.Error()))
	}
}
