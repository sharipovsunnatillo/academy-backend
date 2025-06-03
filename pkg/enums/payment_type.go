package enums

type PaymentType string

const (
	Hourly    PaymentType = "hourly"
	Weekly    PaymentType = "weekly"
	Monthly   PaymentType = "monthly"
	PerLesson PaymentType = "per_lesson"
	Fixed     PaymentType = "fixed"
)
