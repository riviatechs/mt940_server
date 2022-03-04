package models

import (
	"time"

	"github.com/riviatechs/mt940_server/util"
)

type Confirmation struct {
	ID            uint      `gorm:"type:SERIAL PRIMARY KEY"`
	Currency      string    `gorm:"column:currency;type:CHAR(3);NOT NULL"`
	PartyBName    string    `gorm:"column:part_b_name;type:VARCHAR(32);NOT NULL"`
	PartyBAccount string    `gorm:"column:part_b_account;type:VARCHAR(16);NOT NULL"`
	Amount        float32   `gorm:"column:amount;type:NUMERIC;NOT NULL"`
	DateTime      time.Time `gorm:"column:date_time;type:TIMESTAMP;NOT NULL"`
	Narrative     string    `gorm:"column:narrative;type:VARCHAR(64);NOT NULL"`
	PartyAName    string    `gorm:"column:part_a_name;type:VARCHAR(32);NOT NULL"`
	PartyAAccount string    `gorm:"column:part_a_account;type:VARCHAR(16);NOT NULL"`
	Mark          string    `gorm:"column:mark;type:CHAR(1);NOT NULL"`
}

func (Confirmation) TableName() string {
	return util.ConfirmationsTable
}
