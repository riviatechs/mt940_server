package models

import "github.com/riviatechs/mt940_server/util"

type Statement struct {
	ID uint           `gorm:"type:SERIAL PRIMARY KEY"`
	Ob float32        `gorm:"column:ob;type:NUMERIC;NOT NULL"`
	Cb float32        `gorm:"column:cb;type:NUMERIC;NOT NULL"`
	Sl []Confirmation `gorm:"many2many:statement_confirmations;"`
}

func (Statement) TableName() string {
	return util.StatementsTable
}

type StatementConfirmation struct {
	StatementID    uint `gorm:"primaryKey"`
	ConfirmationID uint `gorm:"primaryKey"`
}

func (StatementConfirmation) TableName() string {
	return util.StatementsConfirmationsTable
}
