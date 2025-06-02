package models

type Role struct {
	Name string `gorm:"primaryKey"`
}

type Permission struct {
	Name string `gorm:"primaryKey"`
}
