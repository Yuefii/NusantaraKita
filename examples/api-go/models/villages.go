package models

type Villages struct {
	Code        string `gorm:"column:code"`
	VillageCode string `gorm:"column:village_code"`
	Name        string `gorm:"column:name"`
}
