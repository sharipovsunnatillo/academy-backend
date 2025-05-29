package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sharipov/sunnatillo/academy-backend/internal/api"
	"github.com/sharipov/sunnatillo/academy-backend/internal/database"
	"github.com/sharipov/sunnatillo/academy-backend/internal/models"
	"github.com/sharipov/sunnatillo/academy-backend/internal/repository"
	"github.com/sharipov/sunnatillo/academy-backend/internal/service"
)

func main() {
	fmt.Println("Starting API server...")

	// Initialize database connection
	dbConfig := database.NewDefaultConfig()
	db, err := database.Connect(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto migrate database schemas
	if err := db.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Permission{},
		&models.ProfileImage{},
	); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)

	// Initialize services
	userService := service.NewUserService(userRepo)

	// Initialize handlers
	userHandler := api.NewUserHandler(userService)

	// Register routes
	userHandler.RegisterRoutes()

	// Define a simple handler for the root path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the Academy Backend API!")
	})

	// Start the server
	log.Println("Server listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
