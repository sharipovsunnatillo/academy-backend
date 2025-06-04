package models

type Region struct {
	ID        uint8 `gorm:"primaryKey"`
	Name      string
	Districts []*District `gorm:"foreignKey:RegionID"`
}

type District struct {
	ID       uint16 `gorm:"primaryKey"`
	Name     string
	RegionID uint8
	Region   Region `gorm:"foreignKey:RegionID"`
}

type Address struct {
	RegionID   uint8
	Region     Region `gorm:"foreignKey:RegionID"`
	DistrictID uint16
	District   District `gorm:"foreignKey:DistrictID"`
	Street     string
	House      string
	Apartment  string
	Guide      string
	Longitude  float64
	Latitude   float64
}
