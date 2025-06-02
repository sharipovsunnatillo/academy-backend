package models

type User struct {
	AuditedModel
	FirstName   string
	LastName    string
	MiddleName  string
	Email       string `gorm:"uniqueIndex:users_email_idx"`
	Phone       string `gorm:"uniqueIndex:users_phone_idx"`
	Password    string
	Roles       []*Role       `gorm:"many2many:user_roles;"`
	Permissions []*Permission `gorm:"many2many:user_permissions;"`
}

type TeacherInfo struct {
	ID        uint `gorm:"primaryKey"`
	UserId    uint
	User      User        `gorm:"foreignKey:UserId"`
	Documents []*Document `gorm:"foreignKey:ReferenceId"`
	Subjects  []*Subject  `gorm:"many2many:teacher_subjects;"`
	Branches  []*Branch   `gorm:"many2many:teacher_branches;"`
}
