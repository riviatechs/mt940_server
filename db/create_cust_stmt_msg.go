package db

import (
	"context"
	"fmt"

	"github.com/MalukiMuthusi/logger"
	"github.com/riviatechs/mt940_server/models"
	"go.uber.org/zap"
)

func CreateCustStmtMsg(ctx context.Context, input models.CustStmtMsg) (*int, error) {
	in := &input

	if Db == nil {
		logger.Fatalf("CreateCustStmtMsg", zap.String("db", "db is nil"))
	}

	result := Db.Create(in)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to create a new customer statement message")
	}
	if input.ID == nil {
		return nil, fmt.Errorf("failed to created a new customer statement message")
	}
	id := int(*input.ID)
	return &id, nil
}
