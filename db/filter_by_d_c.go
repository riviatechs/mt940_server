package db

import (
	"context"

	"github.com/riviatechs/mt940_server/models"
)

func FilterByDC(ctx context.Context, input models.DCInput) ([]*models.SlGroups, error) {
	if input.C && input.D {
		return GetStmtLineGroupedByDate(ctx)
	}

	if input.C {
		var sl []*models.Sl
		Db.Where("mark LIKE ?", "C").Find(&sl)

		return GroupStmtLinesByDate(sl)
	}

	if input.D {
		var sl []*models.Sl
		Db.Where("mark LIKE ?", "D").Find(&sl)

		return GroupStmtLinesByDate(sl)
	}

	return GetStmtLineGroupedByDate(ctx)
}
