package model

import "github.com/riviatechs/mt940-server/util"

type Ob struct {
	CustStmtMsgID string `gorm:"column:cus_stmt_msg"`
	Transaction
}

func (Ob) TableName() string {
	return util.ObTable
}
