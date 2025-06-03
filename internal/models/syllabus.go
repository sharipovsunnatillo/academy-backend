package models

import (
	"github.com/sharipov/sunnatillo/academy-backend/pkg/enums"
	"time"
)

type Subject struct {
	ID        uint        `gorm:"primaryKey"`
	Name      string      `gorm:"uniqueIndex:subjects_name_idx"`
	TextBooks []*TextBook `gorm:"foreignKey:SubjectID"`
}

type TextBook struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"uniqueIndex:textbooks_name_idx"`
	Authors   string
	SubjectID uint
	Subject   Subject
}

type TimeSlot struct {
	ID        uint `gorm:"primaryKey"`
	DayOfWeek uint8
	Start     time.Time `gorm:"type:time"`
	End       time.Time `gorm:"type:time"`
}

type Lesson struct {
	ID         uint `gorm:"primaryKey"`
	Name       string
	GroupID    uint
	Group      Group `gorm:"foreignKey:GroupID"`
	SubjectID  uint
	Subject    Subject `gorm:"foreignKey:SubjectID"`
	TimeSlotID uint
	TimeSlot   TimeSlot `gorm:"foreignKey:TimeSlotID"`
	BranchID   uint
	Branch     Branch `gorm:"foreignKey:BranchID"`
	TeacherID  uint
	Teacher    TeacherInfo `gorm:"foreignKey:TeacherID"`
	RoomID     uint
	Room       Room      `gorm:"foreignKey:RoomID"`
	Date       time.Time `gorm:"type:date"`
	Type       enums.LessonType
	Tasks      []*Task `gorm:"foreignKey:LessonID"`
}

type Task struct {
	ID         uint `gorm:"primaryKey"`
	Name       string
	LessonID   uint
	Lesson     Lesson `gorm:"foreignKey:LessonID"`
	DocumentID uint
	Document   Document `gorm:"foreignKey:DocumentID"`
	MaxScore   float32
}
