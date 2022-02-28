package model

import "gorm.io/gorm"

type CustStmtMsg struct {
	gorm.Model
	Trn string
	Rr  *string
	Sn  *string
	Iao *string
}
