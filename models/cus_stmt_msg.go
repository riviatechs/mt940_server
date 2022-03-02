package models

import (
	"time"

	"github.com/riviatechs/mt940_server/util"
	"gorm.io/gorm"
)

type CustStmtMsg struct {
	ID        *uint `gorm:"primarykey"`
	CreatedAt *time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Trn  string  `gorm:"column:trn;not null;unique"`
	Rr   *string `gorm:"column:rr;default null"`
	Ai   Ai      `gorm:"foreignKey:CustStmtMsgID;references:trn"`
	Sn   string  `gorm:"column:sn;not null"`
	Ob   Ob      `gorm:"foreignKey:CustStmtMsgID;references:trn"`
	Sl   []*Sl   `gorm:"foreignKey:CustStmtMsgID;references:trn"`
	Cb   Cb      `gorm:"foreignKey:CustStmtMsgID;references:trn"`
	Cab  *Cab    `gorm:"foreignKey:CustStmtMsgID;references:trn"`
	Fwab []*Fwab `gorm:"foreignKey:CustStmtMsgID;references:trn"`
	Iao  *string `gorm:"column:iao;default null"`
}

func (CustStmtMsg) TableName() string {
	return util.CusStmtMsgTable
}

