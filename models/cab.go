package models

import (
	"github.com/riviatechs/mt940_server/util"
	"gorm.io/gorm"
)

type Cab struct {
	gorm.Model
	CustStmtMsgID string `gorm:"column:cus_stmt_msg;unique"`
	Transaction
}

func (Cab) TableName() string {
	return util.CabTable
}
