package models

import "time"

type AuditedModel struct {
	ID        uint `gorm:"primaryKey"`
	Deleted   bool
	CreatedAt time.Time `gorm:"type:timestamp"`
	CreatedBy uint
	UpdatedAt time.Time `gorm:"type:timestamp"`
	UpdatedBy uint
}
