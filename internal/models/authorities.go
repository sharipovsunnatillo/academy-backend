package models

import "github.com/sharipov/sunnatillo/academy-backend/pkg/enums"

type Role struct {
	Name enums.Role `gorm:"primaryKey"`
}

type Permission struct {
	Name string `gorm:"primaryKey"`
}
