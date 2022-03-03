package models

import (
	"time"

	"github.com/riviatechs/mt940_server/util"
)

type Cab struct {
	CustStmtMsgID uint `gorm:"column:cus_stmt_msg;uniqueIndex;type:INTEGER"`
	Transaction
}

func (Cab) TableName() string {
	return util.CabTable
}

type CabInput struct {
	ID *uint `gorm:"primaryKey"`

	CustStmtMsgID uint `gorm:"column:cus_stmt_msg;uniqueIndex;type:INTEGER"`

	Mark     string    `gorm:"column:mark;type:CHAR NOT NULL"`
	DateY    time.Time `gorm:"column:date_y;type:TIMESTAMP NOT NULL"`
	Currency string    `gorm:"column:currency;type:VARCHAR(3) NOT NULL"`
	Amount   float32   `gorm:"column:amount;type:NUMERIC NOT NULL"`
}

func (CabInput) TableName() string {
	return util.CabTable
}
