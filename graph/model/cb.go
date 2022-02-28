package model

import "github.com/riviatechs/mt940-server/util"

type Cb struct {
	Transaction
}

func (Cb) TableName() string {
	return util.CbTable
}
