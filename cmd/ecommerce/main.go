package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/yourusername/yourproject/config"
	"github.com/yourusername/yourproject/internal/handlers"
	"github.com/yourusername/yourproject/internal/repository"
)

func main() {
	// Load application configurations
	cfg := config.NewConfig()

	// Initialize MongoDB client
	ctx := context.Background()
	userRepo, err := repository.NewUserRepository(ctx, cfg.MongoDBURI, cfg.DatabaseName)
	if err != nil {
		fmt.Printf("Error initializing repository: %v\n", err)
		return
	}

	// Initialize HTTP handlers
	userHandler := handlers.NewUserHandler(userRepo)

	// Routes
	http.HandleFunc("/users", userHandler.GetAllUsers)
	// Add more routes as needed

	// Start server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
