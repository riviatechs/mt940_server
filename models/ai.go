package models

import (
	"time"

	"github.com/riviatechs/mt940_server/util"
	"gorm.io/gorm"
)

type Ai struct {
	ID        *uint `gorm:"primarykey;type: SERIAL PRIMARY KEY"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *gorm.DeletedAt `gorm:"index"`

	CustStmtMsgID uint    `gorm:"column:cus_stmt_msg;uniqueIndex;type:INTEGER"`
	Account       string  `gorm:"column:account;type:VARCHAR(35);NOT NULL"`
	Ic            *string `gorm:"column:ic;type:VARCHAR(11);default:NULL"`
}

func (Ai) TableName() string {
	return util.AiTable
}

type AiInput struct {
	ID            *uint   `gorm:"primaryKey"`
	CustStmtMsgID uint    `gorm:"column:cus_stmt_msg;uniqueIndex;type:INTEGER"`
	Account       string  `gorm:"column:account;type:VARCHAR(35);NOT NULL"`
	Ic            *string `gorm:"column:ic;type:VARCHAR(11);default:NULL"`
}

func (AiInput) TableName() string {
	return util.AiTable
}
