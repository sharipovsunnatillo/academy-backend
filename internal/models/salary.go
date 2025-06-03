package models

import (
	"github.com/sharipov/sunnatillo/academy-backend/pkg/enums"
	"time"
)

type SalaryContract struct {
	AuditedModel
	TeacherID uint
	Teacher   User `gorm:"foreignKey:TeacherID"`
	GroupID   uint
	Group     Group `gorm:"foreignKey:GroupID"`
	Type      enums.PaymentType
	Rate      float64
	Starts    time.Time
	Ends      time.Time
}

type SalaryPayment struct {
	ID               uint `gorm:"primaryKey"`
	SalaryContractID uint
	SalaryContract   SalaryContract `gorm:"foreignKey:SalaryContractID"`
	Amount           float64
	Date             time.Time
}
