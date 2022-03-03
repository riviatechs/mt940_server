package models

import (
	"time"

	"github.com/riviatechs/mt940_server/util"
	"gorm.io/gorm"
)

type CustStmtMsg struct {
	ID        *uint `gorm:"type:SERIAL PRIMARY KEY"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *gorm.DeletedAt `gorm:"index"`

	Trn  string  `gorm:"column:trn;index:idx_cus_stmt_msgs_trn,unique;type:VARCHAR(16);not null"`
	Rr   *string `gorm:"column:rr;type:VARCHAR(16);default:NULL"`
	Ai   Ai      `gorm:"foreignKey:CustStmtMsgID;references:id"`
	Sn   string  `gorm:"column:sn;type:VARCHAR(5);NOT NULL"`
	Ob   Ob      `gorm:"foreignKey:CustStmtMsgID;references:id"`
	Sl   []*Sl   `gorm:"foreignKey:CustStmtMsgID;references:id"`
	Cb   Cb      `gorm:"foreignKey:CustStmtMsgID;references:id"`
	Cab  *Cab    `gorm:"foreignKey:CustStmtMsgID;references:id"`
	Fwab []*Fwab `gorm:"foreignKey:CustStmtMsgID;references:id"`
	Iao  *string `gorm:"column:iao;type:VARCHAR(390);default:NULL"`
}

func (CustStmtMsg) TableName() string {
	return util.CusStmtMsgTable
}

type CustStmtMsgInput struct {
	ID   *uint        `gorm:"primaryKey"`
	Trn  string       `gorm:"column:trn;index:idx_cus_stmt_msgs_trn,unique;type:VARCHAR(16);not null"`
	Rr   *string      `gorm:"column:rr;type:VARCHAR(16);default:NULL"`
	Ai   AiInput      `gorm:"foreignKey:CustStmtMsgID;references:id"`
	Sn   string       `gorm:"column:sn;type:VARCHAR(5):NOT NULL"`
	Ob   ObInput      `gorm:"foreignKey:CustStmtMsgID;references:id"`
	Sl   []*SlInput   `gorm:"foreignKey:CustStmtMsgID;references:id"`
	Cb   CbInput      `gorm:"foreignKey:CustStmtMsgID;references:id"`
	Cab  *CabInput    `gorm:"foreignKey:CustStmtMsgID;references:id"`
	Fwab []*FwabInput `gorm:"foreignKey:CustStmtMsgID;references:id"`
	Iao  *string      `gorm:"column:iao;type:VARCHAR(390);default:NULL"`
}

func (CustStmtMsgInput) TableName() string {
	return util.CusStmtMsgTable
}
