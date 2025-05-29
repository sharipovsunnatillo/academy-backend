package service

import (
	"github.com/sharipov/sunnatillo/academy-backend/internal/models"
	"github.com/sharipov/sunnatillo/academy-backend/internal/repository"
)

// UserService handles business logic for users
type UserService struct {
	repo *repository.UserRepository
}

// NewUserService creates a new user service
func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// Create creates a new user
func (s *UserService) Create(user *models.User) error {
	return s.repo.Create(user)
}

// GetByID retrieves a user by ID
func (s *UserService) GetByID(id uint) (*models.User, error) {
	return s.repo.GetByID(id)
}

// GetByPhone retrieves a user by phone number
func (s *UserService) GetByPhone(phone string) (*models.User, error) {
	return s.repo.GetByPhone(phone)
}

// List retrieves all users with pagination
func (s *UserService) List(page, pageSize int) ([]models.User, int64, error) {
	// Default page and pageSize if not provided
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	
	return s.repo.List(page, pageSize)
}

// Update updates a user
func (s *UserService) Update(user *models.User) error {
	return s.repo.Update(user)
}

// Delete deletes a user by ID
func (s *UserService) Delete(id uint) error {
	return s.repo.Delete(id)
}

// AddRole adds a role to a user
func (s *UserService) AddRole(userID uint, roleID uint) error {
	return s.repo.AddRole(userID, roleID)
}

// RemoveRole removes a role from a user
func (s *UserService) RemoveRole(userID uint, roleID uint) error {
	return s.repo.RemoveRole(userID, roleID)
}

// AddPermission adds a permission to a user
func (s *UserService) AddPermission(userID uint, permissionID uint) error {
	return s.repo.AddPermission(userID, permissionID)
}

// RemovePermission removes a permission from a user
func (s *UserService) RemovePermission(userID uint, permissionID uint) error {
	return s.repo.RemovePermission(userID, permissionID)
}

// AddProfileImage adds a profile image to a user
func (s *UserService) AddProfileImage(profileImage *models.ProfileImage) error {
	return s.repo.AddProfileImage(profileImage)
}

// RemoveProfileImage removes a profile image
func (s *UserService) RemoveProfileImage(imageID uint) error {
	return s.repo.RemoveProfileImage(imageID)
}

// SetMainProfileImage sets a profile image as the main image
func (s *UserService) SetMainProfileImage(userID uint, imageID uint) error {
	return s.repo.SetMainProfileImage(userID, imageID)
}