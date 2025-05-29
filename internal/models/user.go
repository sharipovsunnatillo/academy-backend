package models

import (
	"time"

	"gorm.io/gorm"
)

// Gender type for user gender
type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
	Other  Gender = "other"
)

// User represents a user in the system
type User struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	FirstName   string         `gorm:"size:100;not null" json:"first_name"`
	LastName    string         `gorm:"size:100;not null" json:"last_name"`
	MiddleName  string         `gorm:"size:100" json:"middle_name"`
	Phone       string         `gorm:"size:20;uniqueIndex" json:"phone"`
	Roles       []Role         `gorm:"many2many:user_roles;" json:"roles"`
	Permissions []Permission   `gorm:"many2many:user_permissions;" json:"permissions"`
	Birthday    *time.Time     `json:"birthday"`
	Gender      Gender         `gorm:"size:10" json:"gender"`
	ProfileImages []ProfileImage `gorm:"foreignKey:UserID" json:"profile_images"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// Role represents a role in the system
type Role struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"size:100;uniqueIndex;not null" json:"name"`
	Description string         `gorm:"size:255" json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// Permission represents a permission in the system
type Permission struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"size:100;uniqueIndex;not null" json:"name"`
	Description string         `gorm:"size:255" json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// ProfileImage represents a user profile image
type ProfileImage struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	UserID    uint           `gorm:"not null" json:"user_id"`
	URL       string         `gorm:"size:255;not null" json:"url"`
	IsMain    bool           `gorm:"default:false" json:"is_main"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// TableName overrides the table name
func (User) TableName() string {
	return "users"
}

// TableName overrides the table name
func (Role) TableName() string {
	return "roles"
}

// TableName overrides the table name
func (Permission) TableName() string {
	return "permissions"
}

// TableName overrides the table name
func (ProfileImage) TableName() string {
	return "profile_images"
}