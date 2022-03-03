package db

import (
	"context"

	"github.com/riviatechs/mt940_server/models"
)

func GetStmtLinesFilterByPeriod(ctx context.Context, amount models.Amount) ([]*models.SlGroups, error) {
	var sl []*models.Sl
	Db.Where("amount BETWEEN ? AND ?", amount.Lower, amount.Upper).Find(&sl)

	return GroupStmtLinesByDate(sl)
}
