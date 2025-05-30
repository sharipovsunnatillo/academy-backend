package enums

type Role string

const (
	SUPER_ADMIN Role = "SUPER_ADMIN"
	ADMIN       Role = "ADMIN"
	TEACHER     Role = "TEACHER"
	STUDENT     Role = "STUDENT"
	GUEST       Role = "GUEST"
	PARENT      Role = "PARENT"
)
