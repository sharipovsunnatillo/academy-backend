package models

import (
	"database/sql"
	"github.com/sharipov/sunnatillo/academy-backend/pkg/enums"
)

type Role struct {
	Name string `gorm:"primaryKey;type:varchar(50)"`
}

type Permission struct {
	Name string `gorm:"primaryKey;type:varchar(255)"`
}

type User struct {
	AuditedModel
	Active         bool
	ActivationCode sql.NullString         `gorm:"type:varchar(6)"`
	ActivationDate sql.NullTime           `gorm:"type:timestamp"`
	ResetCode      sql.NullString         `gorm:"type:varchar(6)"`
	ResetDate      sql.NullTime           `gorm:"type:timestamp"`
	FirstName      string                 `gorm:"type:varchar(50)"`
	LastName       string                 `gorm:"type:varchar(50)"`
	MiddleName     string                 `gorm:"type:varchar(50)"`
	Gender         sql.Null[enums.Gender] `gorm:"type:varchar(15)"`
	Birthday       sql.NullTime           `gorm:"type:date"`
	Phone          string                 `gorm:"unique;type:varchar(20)"`
	Password       string                 `gorm:"type:varchar(255)"`
	Roles          []*Role                `gorm:"many2many:user_roles;"`
	Permissions    []*Permission          `gorm:"many2many:user_permissions;"`
}
