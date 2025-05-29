package repository

import (
	"errors"

	"github.com/sharipov/sunnatillo/academy-backend/internal/models"
	"gorm.io/gorm"
)

// UserRepository handles database operations for users
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create creates a new user
func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

// GetByID retrieves a user by ID
func (r *UserRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	result := r.db.Preload("Roles").Preload("Permissions").Preload("ProfileImages").First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}
	return &user, nil
}

// GetByPhone retrieves a user by phone number
func (r *UserRepository) GetByPhone(phone string) (*models.User, error) {
	var user models.User
	result := r.db.Preload("Roles").Preload("Permissions").Preload("ProfileImages").Where("phone = ?", phone).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}
	return &user, nil
}

// List retrieves all users with pagination
func (r *UserRepository) List(page, pageSize int) ([]models.User, int64, error) {
	var users []models.User
	var count int64

	// Count total records
	if err := r.db.Model(&models.User{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated records
	offset := (page - 1) * pageSize
	result := r.db.Preload("Roles").Preload("Permissions").Preload("ProfileImages").
		Offset(offset).Limit(pageSize).Find(&users)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return users, count, nil
}

// Update updates a user
func (r *UserRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

// Delete deletes a user by ID
func (r *UserRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}

// AddRole adds a role to a user
func (r *UserRepository) AddRole(userID uint, roleID uint) error {
	var user models.User
	if err := r.db.First(&user, userID).Error; err != nil {
		return err
	}

	var role models.Role
	if err := r.db.First(&role, roleID).Error; err != nil {
		return err
	}

	return r.db.Model(&user).Association("Roles").Append(&role)
}

// RemoveRole removes a role from a user
func (r *UserRepository) RemoveRole(userID uint, roleID uint) error {
	var user models.User
	if err := r.db.First(&user, userID).Error; err != nil {
		return err
	}

	var role models.Role
	if err := r.db.First(&role, roleID).Error; err != nil {
		return err
	}

	return r.db.Model(&user).Association("Roles").Delete(&role)
}

// AddPermission adds a permission to a user
func (r *UserRepository) AddPermission(userID uint, permissionID uint) error {
	var user models.User
	if err := r.db.First(&user, userID).Error; err != nil {
		return err
	}

	var permission models.Permission
	if err := r.db.First(&permission, permissionID).Error; err != nil {
		return err
	}

	return r.db.Model(&user).Association("Permissions").Append(&permission)
}

// RemovePermission removes a permission from a user
func (r *UserRepository) RemovePermission(userID uint, permissionID uint) error {
	var user models.User
	if err := r.db.First(&user, userID).Error; err != nil {
		return err
	}

	var permission models.Permission
	if err := r.db.First(&permission, permissionID).Error; err != nil {
		return err
	}

	return r.db.Model(&user).Association("Permissions").Delete(&permission)
}

// AddProfileImage adds a profile image to a user
func (r *UserRepository) AddProfileImage(profileImage *models.ProfileImage) error {
	return r.db.Create(profileImage).Error
}

// RemoveProfileImage removes a profile image
func (r *UserRepository) RemoveProfileImage(imageID uint) error {
	return r.db.Delete(&models.ProfileImage{}, imageID).Error
}

// SetMainProfileImage sets a profile image as the main image
func (r *UserRepository) SetMainProfileImage(userID uint, imageID uint) error {
	// First, set all images for this user to not main
	if err := r.db.Model(&models.ProfileImage{}).Where("user_id = ?", userID).Update("is_main", false).Error; err != nil {
		return err
	}

	// Then set the specified image as main
	return r.db.Model(&models.ProfileImage{}).Where("id = ? AND user_id = ?", imageID, userID).Update("is_main", true).Error
}