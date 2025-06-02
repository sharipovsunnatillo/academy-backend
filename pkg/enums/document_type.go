package enums

type DocumentType string

const (
	Passport      DocumentType = "passport"
	IDCard        DocumentType = "id_card"
	DriverLicense DocumentType = "driver_license"
	Certificate   DocumentType = "certificate"
	Other         DocumentType = "other"
)
