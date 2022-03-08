package db

import (
	"context"
	"fmt"

	"github.com/MalukiMuthusi/logger"
	"github.com/riviatechs/mt940_server/models"
	"go.uber.org/zap"
)

func Statements(ctx context.Context) ([]*models.Confirmation, error) {
	var confirmations []*models.Confirmation

	Db.Order("date_time desc").Find(&confirmations)

	if confirmations == nil {
		return nil, fmt.Errorf("failed to get statements")
	}

	return confirmations, nil
}

func GetStatement(ctx context.Context, id uint) (*models.Confirmation, error) {
	var confirmation = &models.Confirmation{}
	logger.Info("GetStatement", zap.Uint("id", id))
	Db.First(confirmation, id)

	if confirmation == nil {
		return nil, fmt.Errorf("record not found")
	}

	return confirmation, nil
}
