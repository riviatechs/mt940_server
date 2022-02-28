package model

import (
	"github.com/riviatechs/mt940-server/util"
	"gorm.io/gorm"
)

type CustStmtMsg struct {
	gorm.Model
	Trn  string  `gorm:"column:trn"`
	Rr   *string `gorm:"column:rr"`
	Ai   Ai      `gorm:"column:ai;foreignKey:CustStmtMsgID"`
	Sn   string  `gorm:"column:sn"`
	Ob   Ob      `gorm:"column:ob;foreignKey:CustStmtMsgID"`
	Sl   []Sl    `gorm:"column:sl;foreignKey:CustStmtMsgID"`
	Cb   Cb      `gorm:"column:cb;foreignKey:CustStmtMsgID"`
	Cab  *Cab    `gorm:"column:cab;foreignKey:CustStmtMsgID"`
	Fwab []Fwab  `gorm:"column:fwab;foreignKey:CustStmtMsgID"`
	Iao  *string `gorm:"column:iao"`
}

func (CustStmtMsg) TableName() string {
	return util.CusStmtMsgTable
}
