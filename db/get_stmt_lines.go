package db

import (
	"context"

	"github.com/riviatechs/mt940_server/models"
)

func GetStatementLines(ctx context.Context) ([]*models.Sl, error) {
	var sl []*models.Sl

	Db.Find(&sl)

	return sl, nil
}
