package model

import "github.com/riviatechs/mt940-server/util"

type Cb struct {
	CustStmtMsgID string `gorm:"column:cus_stmt_msg;unique"`
	Transaction
}

func (Cb) TableName() string {
	return util.CbTable
}
