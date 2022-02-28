package model

import (
	"github.com/riviatechs/mt940-server/util"
	"gorm.io/gorm"
)

type Sl struct {
	gorm.Model
	CustStmtMsgID string  `gorm:"column:cus_stmt_msg"`
	ValueDate     *string `gorm:"column:value_date"`
	EntryDate     *string `gorm:"column:entry_date"`
	Mark          string  `gorm:"column:mark"`
	FundsCode     *string `gorm:"column:funds_code"`
	Amount        float32 `gorm:"column:amount"`
	Ttic          *string `gorm:"column:tt_ic"`
	RefOwner      *string `gorm:"column:ref_owner"`
	RefAsi        *string `gorm:"column:ref_asi"`
	Supp          *string `gorm:"column:supp"`
	Iao           *string `gorm:"column:iao"`
}

func (Sl) TableName() string {
	return util.SlTable
}
