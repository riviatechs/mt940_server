package model

import "github.com/riviatechs/mt940-server/util"

type Cab struct {
	Transaction
}

func (Cab) TableName() string {
	return util.CabTable
}
