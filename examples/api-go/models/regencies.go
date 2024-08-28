package models

type Regencies struct {
	Code         string `gorm:"column:code"`
	ProvinceCode string `gorm:"column:province_code"`
	Name         string `gorm:"column:name"`
}
