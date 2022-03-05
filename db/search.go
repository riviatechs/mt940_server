package db

import (
	"context"
	"fmt"

	"github.com/riviatechs/mt940_server/models"
)

func Search(ctx context.Context, input string) ([]*models.ConfGroup, error) {
	db := Db.Where("searchable @@ websearch_to_tsquery(?)", input)

	var confirmations []*models.Confirmation

	db.Order("date_time desc").Find(&confirmations)

	if confirmations == nil {
		return nil, fmt.Errorf("failed to get statements")
	}

	return GroupStmtsByDate(confirmations)
}
