package models

import (
	"github.com/sharipov/sunnatillo/academy-backend/pkg/enums"
	"time"
)

type Group struct {
	AuditedModel
	Name      string
	BranchID  uint
	Branch    Branch `gorm:"foreignKey:BranchID"`
	TeacherID uint
	Teacher   TeacherInfo `gorm:"foreignKey:TeacherID"`
	SubjectID uint
	Subject   Subject `gorm:"foreignKey:SubjectID"`
	RoomID    uint
	Room      Room        `gorm:"foreignKey:RoomID"`
	TimeSlots []*TimeSlot `gorm:"many2many:group_timeslots;"`
	Students  []*User     `gorm:"many2many:group_students;"`
	Started   time.Time
	Ended     time.Time
	Period    uint16
	Duration  enums.TimeDuration
}
