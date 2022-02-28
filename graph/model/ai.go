package model

import (
	"github.com/riviatechs/mt940-server/util"
	"gorm.io/gorm"
)

type Ai struct {
	gorm.Model
	CustStmtMsgID string  `gorm:"column:cus_stmt_msg"`
	Account       string  `gorm:"column:account"`
	Ic            *string `gorm:"column:ic"`
}

func (Ai) TableName() string {
	return util.AiTable
}
