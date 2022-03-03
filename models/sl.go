package models

import (
	"time"

	"github.com/riviatechs/mt940_server/util"
	"gorm.io/gorm"
)

type Sl struct {
	ID        *uint `gorm:"primarykey;type:SERIAL PRIMARY KEY"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *gorm.DeletedAt `gorm:"index"`

	CustStmtMsgID uint       `gorm:"column:cus_stmt_msg;type:INTEGER REFERENCES cus_stmt_msgs(id)"`
	ValueDate     time.Time  `gorm:"column:value_date;type:TIMESTAMP NOT NULL"`
	EntryDate     *time.Time `gorm:"column:entry_date;type:TIMESTAMP DEFAULT NULL"`
	Mark          string     `gorm:"column:mark;type:VARCHAR(2) NOT NULL"`
	FundsCode     *string    `gorm:"column:funds_code;type:CHAR(1) DEFAULT NULL"`
	Amount        float32    `gorm:"column:amount;type:NUMERIC(15) NOT NULL"`
	Ttic          *string    `gorm:"column:tt_ic;type:VARCHAR(4) DEFAULT NULL"`
	RefOwner      *string    `gorm:"column:ref_owner;type:VARCHAR(16)  DEFAULT NULL"`
	RefAsi        *string    `gorm:"column:ref_asi;type:VARCHAR(16) DEFAULT NULL"`
	Supp          *string    `gorm:"column:supp;type:VARCHAR(34) DEFAULT NULL"`
	Iao           *string    `gorm:"column:iao;type:VARCHAR(390) DEFAULT NULL"`
}

func (Sl) TableName() string {
	return util.SlTable
}

type SlInput struct {
	ID *uint `gorm:"primaryKey"`

	CustStmtMsgID uint       `gorm:"column:cus_stmt_msg;type:INTEGER REFERENCES cus_stmt_msgs(id)"`
	ValueDate     time.Time  `gorm:"column:value_date;type:TIMESTAMP NOT NULL"`
	EntryDate     *time.Time `gorm:"column:entry_date;type:TIMESTAMP DEFAULT NULL"`
	Mark          string     `gorm:"column:mark;type:VARCHAR(2) NOT NULL"`
	FundsCode     *string    `gorm:"column:funds_code;type:CHAR(1) DEFAULT NULL"`
	Amount        float32    `gorm:"column:amount;type:NUMERIC(15) NOT NULL"`
	Ttic          *string    `gorm:"column:tt_ic;type:VARCHAR(4) DEFAULT NULL"`
	RefOwner      *string    `gorm:"column:ref_owner;type:VARCHAR(16)  DEFAULT NULL"`
	RefAsi        *string    `gorm:"column:ref_asi;type:VARCHAR(16) DEFAULT NULL"`
	Supp          *string    `gorm:"column:supp;type:VARCHAR(34) DEFAULT NULL"`
	Iao           *string    `gorm:"column:iao;type:VARCHAR(390) DEFAULT NULL"`
}

func (SlInput) TableName() string {
	return util.SlTable
}
