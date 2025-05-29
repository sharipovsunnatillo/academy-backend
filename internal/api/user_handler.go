package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/sharipov/sunnatillo/academy-backend/internal/models"
	"github.com/sharipov/sunnatillo/academy-backend/internal/service"
)

// UserHandler handles HTTP requests for users
type UserHandler struct {
	userService *service.UserService
}

// NewUserHandler creates a new user handler
func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// RegisterRoutes registers the user routes
func (h *UserHandler) RegisterRoutes() {
	http.HandleFunc("/api/users", h.handleUsers)
	http.HandleFunc("/api/users/", h.handleUserByID)
}

// handleUsers handles requests to /api/users
func (h *UserHandler) handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.listUsers(w, r)
	case http.MethodPost:
		h.createUser(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// handleUserByID handles requests to /api/users/{id}
func (h *UserHandler) handleUserByID(w http.ResponseWriter, r *http.Request) {
	// Extract ID from URL
	idStr := r.URL.Path[len("/api/users/"):]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.getUserByID(w, r, uint(id))
	case http.MethodPut:
		h.updateUser(w, r, uint(id))
	case http.MethodDelete:
		h.deleteUser(w, r, uint(id))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// listUsers lists all users
func (h *UserHandler) listUsers(w http.ResponseWriter, r *http.Request) {
	// Parse pagination parameters
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("pageSize")

	page := 1
	pageSize := 10

	if pageStr != "" {
		pageVal, err := strconv.Atoi(pageStr)
		if err == nil && pageVal > 0 {
			page = pageVal
		}
	}

	if pageSizeStr != "" {
		pageSizeVal, err := strconv.Atoi(pageSizeStr)
		if err == nil && pageSizeVal > 0 {
			pageSize = pageSizeVal
		}
	}

	// Get users from service
	users, total, err := h.userService.List(page, pageSize)
	if err != nil {
		http.Error(w, "Failed to retrieve users: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Prepare response
	response := struct {
		Users []models.User `json:"users"`
		Total int64         `json:"total"`
		Page  int           `json:"page"`
		Size  int           `json:"size"`
	}{
		Users: users,
		Total: total,
		Page:  page,
		Size:  pageSize,
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

// createUser creates a new user
func (h *UserHandler) createUser(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Create user
	if err := h.userService.Create(&user); err != nil {
		http.Error(w, "Failed to create user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

// getUserByID gets a user by ID
func (h *UserHandler) getUserByID(w http.ResponseWriter, r *http.Request, id uint) {
	// Get user from service
	user, err := h.userService.GetByID(id)
	if err != nil {
		http.Error(w, "Failed to retrieve user: "+err.Error(), http.StatusNotFound)
		return
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

// updateUser updates a user
func (h *UserHandler) updateUser(w http.ResponseWriter, r *http.Request, id uint) {
	// Get existing user
	existingUser, err := h.userService.GetByID(id)
	if err != nil {
		http.Error(w, "User not found: "+err.Error(), http.StatusNotFound)
		return
	}

	// Parse request body
	var updatedUser models.User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Update user fields
	updatedUser.ID = id
	updatedUser.CreatedAt = existingUser.CreatedAt
	updatedUser.UpdatedAt = time.Now()

	// Update user
	if err := h.userService.Update(&updatedUser); err != nil {
		http.Error(w, "Failed to update user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(updatedUser); err != nil {
		http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

// deleteUser deletes a user
func (h *UserHandler) deleteUser(w http.ResponseWriter, r *http.Request, id uint) {
	// Delete user
	if err := h.userService.Delete(id); err != nil {
		http.Error(w, "Failed to delete user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Send response
	w.WriteHeader(http.StatusNoContent)
}