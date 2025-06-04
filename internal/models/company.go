package models

type TrainingCenter struct {
	AuditedModel
	Name string
	Address
	Branches []*Branch `gorm:"foreignKey:CenterID"`
}
type Branch struct {
	AuditedModel
	Name     string
	CenterID uint
	Center   TrainingCenter `gorm:"foreignKey:CenterID"`
	Address
}

type Room struct {
	ID       uint `gorm:"primaryKey"`
	Number   string
	BranchID uint
	Branch   *Branch `gorm:"foreignKey:BranchID"`
}
