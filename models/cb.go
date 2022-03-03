package models

import (
	"time"

	"github.com/riviatechs/mt940_server/util"
)

type Cb struct {
	CustStmtMsgID uint `gorm:"column:cus_stmt_msg;type:INTEGER UNIQUE REFERENCES cus_stmt_msgs(id)"`
	Transaction
}

func (Cb) TableName() string {
	return util.CbTable
}

type CbInput struct {
	ID *uint `gorm:"primaryKey"`

	CustStmtMsgID uint `gorm:"column:cus_stmt_msg;type:INTEGER UNIQUE REFERENCES cus_stmt_msgs(id)"`

	Mark     string    `gorm:"column:mark;type:CHAR NOT NULL"`
	DateY    time.Time `gorm:"column:date_y;type:TIMESTAMP NOT NULL"`
	Currency string    `gorm:"column:currency;type:VARCHAR(3) NOT NULL"`
	Amount   float32   `gorm:"column:amount;type:NUMERIC NOT NULL"`
}

func (CbInput) TableName() string {
	return util.CbTable
}
