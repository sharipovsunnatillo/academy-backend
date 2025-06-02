package models

import "time"

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
}
