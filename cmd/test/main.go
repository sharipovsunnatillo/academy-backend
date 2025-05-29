package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/sharipov/sunnatillo/academy-backend/internal/models"
)

func main() {
	baseURL := "http://localhost:8080/api/users"

	// Test creating a user
	fmt.Println("Testing user creation...")

	// Create a sample user
	birthday := time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)
	newUser := models.User{
		FirstName:  "John",
		LastName:   "Doe",
		MiddleName: "Smith",
		Phone:      "+1234567890",
		Birthday:   &birthday,
		Gender:     models.Male,
	}

	// Convert user to JSON
	userJSON, err := json.Marshal(newUser)
	if err != nil {
		log.Fatalf("Failed to marshal user: %v", err)
	}

	// Send POST request
	resp, err := http.Post(baseURL, "application/json", bytes.NewBuffer(userJSON))
	if err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}
	defer resp.Body.Close()

	// Check response
	if resp.StatusCode != http.StatusCreated {
		log.Fatalf("Expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
	}

	// Decode created user
	var createdUser models.User
	if err := json.NewDecoder(resp.Body).Decode(&createdUser); err != nil {
		log.Fatalf("Failed to decode response: %v", err)
	}

	fmt.Printf("User created successfully with ID: %d\n", createdUser.ID)

	// Test getting a user by ID
	fmt.Println("\nTesting get user by ID...")

	// Send GET request
	getUserURL := fmt.Sprintf("%s/%d", baseURL, createdUser.ID)
	resp, err = http.Get(getUserURL)
	if err != nil {
		log.Fatalf("Failed to get user: %v", err)
	}
	defer resp.Body.Close()

	// Check response
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	// Decode user
	var retrievedUser models.User
	if err := json.NewDecoder(resp.Body).Decode(&retrievedUser); err != nil {
		log.Fatalf("Failed to decode response: %v", err)
	}

	fmt.Printf("User retrieved successfully: %s %s\n", retrievedUser.FirstName, retrievedUser.LastName)

	// Test updating a user
	fmt.Println("\nTesting user update...")

	// Update user
	retrievedUser.FirstName = "Jane"
	retrievedUser.LastName = "Smith"

	// Convert updated user to JSON
	updatedUserJSON, err := json.Marshal(retrievedUser)
	if err != nil {
		log.Fatalf("Failed to marshal updated user: %v", err)
	}

	// Create PUT request
	req, err := http.NewRequest(http.MethodPut, getUserURL, bytes.NewBuffer(updatedUserJSON))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Send PUT request
	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		log.Fatalf("Failed to update user: %v", err)
	}
	defer resp.Body.Close()

	// Check response
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	// Decode updated user
	var updatedUser models.User
	if err := json.NewDecoder(resp.Body).Decode(&updatedUser); err != nil {
		log.Fatalf("Failed to decode response: %v", err)
	}

	fmt.Printf("User updated successfully: %s %s\n", updatedUser.FirstName, updatedUser.LastName)

	// Test listing users
	fmt.Println("\nTesting list users...")

	// Send GET request
	resp, err = http.Get(baseURL)
	if err != nil {
		log.Fatalf("Failed to list users: %v", err)
	}
	defer resp.Body.Close()

	// Check response
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	// Decode users list
	var usersResponse struct {
		Users []models.User `json:"users"`
		Total int64         `json:"total"`
		Page  int           `json:"page"`
		Size  int           `json:"size"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&usersResponse); err != nil {
		log.Fatalf("Failed to decode response: %v", err)
	}

	fmt.Printf("Retrieved %d users out of %d total\n", len(usersResponse.Users), usersResponse.Total)

	// Test deleting a user
	fmt.Println("\nTesting user deletion...")

	// Create DELETE request
	req, err = http.NewRequest(http.MethodDelete, getUserURL, nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	// Send DELETE request
	resp, err = client.Do(req)
	if err != nil {
		log.Fatalf("Failed to delete user: %v", err)
	}
	defer resp.Body.Close()

	// Check response
	if resp.StatusCode != http.StatusNoContent {
		log.Fatalf("Expected status code %d, got %d", http.StatusNoContent, resp.StatusCode)
	}

	fmt.Println("User deleted successfully")

	fmt.Println("\nAll tests passed successfully!")
}
