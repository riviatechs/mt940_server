package model

import "gorm.io/gorm"

type CustStmtMsg struct {
	gorm.Model
	Trn  string
	Rr   *string
	Ai   Ai
	Sn   string
	Ob   Transaction
	Sl   []*Sl
	Cb   Transaction
	Cab  *Transaction
	Fwab []*Transaction
	Iao  *string
}
