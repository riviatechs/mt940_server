package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID        *uint `gorm:"primarykey"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *gorm.DeletedAt `gorm:"index"`

	Mark     string    `gorm:"column:mark;not null"`
	DateY    time.Time `gorm:"column:date_y;not null"`
	Currency string    `gorm:"column:currency;not null"`
	Amount   float32   `gorm:"column:amount;not null"`
}
