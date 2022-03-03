package db

import (
	"context"

	"github.com/riviatechs/mt940_server/models"
)

func GetStmtLineGroupedByDate(ctx context.Context) ([]*models.SlGroups, error) {
	var sl []*models.Sl
	Db.Find(&sl)

	return GroupStmtLinesByDate(sl)
}
