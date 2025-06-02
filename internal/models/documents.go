package models

import "github.com/sharipov/sunnatillo/academy-backend/pkg/enums"

type Document struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Url         string
	Type        enums.DocumentType
	ReferenceId uint
}
