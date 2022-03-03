package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID        *uint           `gorm:"primarykey;type:SERIAL PRIMARY KEY"`
	CreatedAt *time.Time      `gorm:"type:TIMESTAMP;default:null"`
	UpdatedAt *time.Time      `gorm:"type:TIMESTAMP;default:null"`
	DeletedAt *gorm.DeletedAt `gorm:"index;type:TIMESTAMP DEFAULT NULL"`

	Mark     string    `gorm:"column:mark;type:CHAR;NOT NULL"`
	DateY    time.Time `gorm:"column:date_y;type:TIMESTAMP;NOT NULL"`
	Currency string    `gorm:"column:currency;type:VARCHAR(3);NOT NULL"`
	Amount   float32   `gorm:"column:amount;type:NUMERIC;NOT NULL"`
}

type TransactionInput struct {
	ID       *uint     `gorm:"primaryKey"`
	Mark     string    `gorm:"column:mark;type:CHAR;NOT NULL"`
	DateY    time.Time `gorm:"column:date_y;type:TIMESTAMP;NOT NULL"`
	Currency string    `gorm:"column:currency;type:VARCHAR(3);NOT NULL"`
	Amount   float32   `gorm:"column:amount;type:NUMERIC;NOT NULL"`
}
