package models

import (
	"time"

	"github.com/riviatechs/mt940_server/util"
	"gorm.io/gorm"
)

type Ai struct {
	ID        *uint `gorm:"primarykey"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *gorm.DeletedAt `gorm:"index"`

	CustStmtMsgID string  `gorm:"column:cus_stmt_msg;unique"`
	Account       string  `gorm:"column:account;not null"`
	Ic            *string `gorm:"column:ic;default null"`
}

func (Ai) TableName() string {
	return util.AiTable
}
