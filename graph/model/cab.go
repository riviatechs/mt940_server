package model

import "github.com/riviatechs/mt940-server/util"

type Cab struct {
	CustStmtMsgID string `gorm:"column:cus_stmt_msg;unique"`
	Transaction
}

func (Cab) TableName() string {
	return util.CabTable
}
