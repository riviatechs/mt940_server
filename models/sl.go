package models

import (
	"time"

	"github.com/riviatechs/mt940_server/util"
	"gorm.io/gorm"
)

type Sl struct {
	ID        *uint `gorm:"primarykey"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *gorm.DeletedAt `gorm:"index"`

	CustStmtMsgID string     `gorm:"column:cus_stmt_msg"`
	ValueDate     time.Time  `gorm:"column:value_date;not null"`
	EntryDate     *time.Time `gorm:"column:entry_date;default null"`
	Mark          string     `gorm:"column:mark;default null"`
	FundsCode     *string    `gorm:"column:funds_code;default null"`
	Amount        float32    `gorm:"column:amount;not null"`
	Ttic          *string    `gorm:"column:tt_ic;default null"`
	RefOwner      *string    `gorm:"column:ref_owner;default null"`
	RefAsi        *string    `gorm:"column:ref_asi;default null"`
	Supp          *string    `gorm:"column:supp;default null"`
	Iao           *string    `gorm:"column:iao;default null"`
}

func (Sl) TableName() string {
	return util.SlTable
}
