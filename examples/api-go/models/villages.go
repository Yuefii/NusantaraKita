package models

type Villages struct {
	Code         string    `gorm:"column:code"`
	DistrictCode string    `gorm:"column:district_code"`
	Name         string    `gorm:"column:name"`
	District     Districts `gorm:"foreignkey:DistrictCode;references:Code"`
}
