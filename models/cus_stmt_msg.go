package models

import (
	"github.com/riviatechs/mt940_server/util"
	"gorm.io/gorm"
)

type CustStmtMsg struct {
	gorm.Model
	Trn  string  `gorm:"column:trn;not null;unique"`
	Rr   *string `gorm:"column:rr;default null"`
	Ai   Ai      `gorm:"column:ai;foreignKey:CustStmtMsgID;references:Trn"`
	Sn   string  `gorm:"column:sn;not null"`
	Ob   Ob      `gorm:"column:ob;foreignKey:CustStmtMsgID;references:Trn"`
	Sl   []Sl    `gorm:"column:sl;foreignKey:CustStmtMsgID;references:Trn"`
	Cb   Cb      `gorm:"column:cb;foreignKey:CustStmtMsgID;references:Trn"`
	Cab  *Cab    `gorm:"column:cab;foreignKey:CustStmtMsgID;references:Trn"`
	Fwab []Fwab  `gorm:"column:fwab;foreignKey:CustStmtMsgID;references:Trn"`
	Iao  *string `gorm:"column:iao;default null"`
}

func (CustStmtMsg) TableName() string {
	return util.CusStmtMsgTable
}
