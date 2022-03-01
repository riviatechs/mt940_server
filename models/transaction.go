package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Mark     string  `gorm:"column:mark;not null"`
	DateY    string  `gorm:"column:date_y;not null"`
	Currency string  `gorm:"column:currency;not null"`
	Amount   float32 `gorm:"column:amount;not null"`
}
