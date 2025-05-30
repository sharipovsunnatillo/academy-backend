package models

import "time"

type AuditedModel struct {
	ID        uint      `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"type:timestamp"`
	CreatedBy uint
	UpdatedAt time.Time `gorm:"type:timestamp"`
	UpdatedBy uint
}
