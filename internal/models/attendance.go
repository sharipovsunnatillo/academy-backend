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

type Grade struct {
	ID        uint
	StudentID uint
	Student   User `gorm:"foreignKey:StudentID"`
	LessonID  uint
	Lesson    Lesson `gorm:"foreignKey:LessonID"`
	TaskID    uint
	Task      Task `gorm:"foreignKey:TaskID"`
	Score     float32
}
