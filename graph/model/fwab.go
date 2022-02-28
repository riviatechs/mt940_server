package model

import "github.com/riviatechs/mt940-server/util"

type Fwab struct {
	CustStmtMsgID string `gorm:"column:cus_stmt_msg"`
	Transaction
}

func (Fwab) TableName() string {
	return util.FwabTable
}
