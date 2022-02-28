package model

import "github.com/riviatechs/mt940-server/util"

type Fwab struct {
	Transaction
}

func (Fwab) TableName() string {
	return util.FwabTable
}
