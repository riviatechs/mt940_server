package model

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	CustStmtMsgID string  `gorm:"column:cus_stmt_msg"`
	Mark          string  `gorm:"column:mark"`
	DateY         string  `gorm:"column:date_y"`
	Currency      string  `gorm:"column:currency"`
	Amount        float32 `gorm:"column:amount"`
}
