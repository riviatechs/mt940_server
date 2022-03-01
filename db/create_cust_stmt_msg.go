package db

import (
	"context"
	"fmt"

	"github.com/riviatechs/mt940_server/models"
)

func CreateCustStmtMsg(ctx context.Context, input models.CustStmtMsg) (*int, error) {
	in := input
	result := Db.Create(&in)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to create a new customer statement message")
	}
	if input.ID == nil {
		return nil, fmt.Errorf("failed to created a new customer statement message")
	}
	id := int(*input.ID)
	return &id, nil
}
