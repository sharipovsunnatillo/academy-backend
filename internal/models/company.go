package models

type TrainingCenter struct {
	AuditedModel
	Name string
	Address
	Branches []*Branch `gorm:"foreignKey:CenterId"`
}
type Branch struct {
	AuditedModel
	Name     string
	CenterId uint
	Center   TrainingCenter
	Address
}

type Room struct {
	Number string
	Branch *Branch
}
