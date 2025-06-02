package models

import (
	"github.com/sharipov/sunnatillo/academy-backend/pkg/enums"
	"time"
)

type Group struct {
	AuditedModel
	Name       string
	TeacherId  uint
	Teacher    TeacherInfo `gorm:"foreignKey:TeacherId"`
	SubjectId  uint
	Subject    Subject `gorm:"foreignKey:SubjectId"`
	RoomId     uint
	Room       Room `gorm:"foreignKey:RoomId"`
	TimeSlotId uint
	TimeSlot   TimeSlot `gorm:"foreignKey:TimeSlotId"`
	Students   []*User  `gorm:"many2many:group_students;"`
	Started    time.Time
	Ended      time.Time
	Period     uint16
	Duration   enums.TimeDuration
}
