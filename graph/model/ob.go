package model

import "github.com/riviatechs/mt940-server/util"

type Ob struct {
	Transaction
}

func (Ob) TableName() string {
	return util.ObTable
}
