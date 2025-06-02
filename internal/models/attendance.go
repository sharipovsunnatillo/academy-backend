package models

import "github.com/sharipov/sunnatillo/academy-backend/pkg/enums"

type Attendance struct {
	ID        uint
	LessonID  uint
	Lesson    Lesson `gorm:"foreignKey:LessonID"`
	StudentID uint
	Student   User `gorm:"foreignKey:StudentID"`
	Status    enums.AttendanceStatus
}
