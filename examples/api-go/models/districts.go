package models

type Districts struct {
	Code        string    `gorm:"column:code"`
	RegencyCode string    `gorm:"column:regency_code"`
	Name        string    `gorm:"column:name"`
	Regency     Regencies `gorm:"foreignkey:RegencyCode;references:Code"`
}
