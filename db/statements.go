package db

import (
	"context"
	"fmt"

	"github.com/riviatechs/mt940_server/models"
)

func Statements(ctx context.Context) ([]*models.Confirmation, error) {
	var confirmations []*models.Confirmation

	Db.Order("date_time desc").Find(&confirmations)

	if confirmations == nil {
		return nil, fmt.Errorf("failed to get statements")
	}

	return confirmations, nil
}
