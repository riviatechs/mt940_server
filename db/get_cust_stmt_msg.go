package db

import (
	"context"
	"fmt"

	"github.com/riviatechs/mt940_server/models"
)

func GetCustStmtMsg(ctx context.Context, id int) (*models.CustStmtMsg, error) {
	custStmtMsg := &models.CustStmtMsg{}

	Db.Preload("Ai").Preload("Ob").Preload("Sl").Preload("Cb").Preload("Cab").Preload("Fwab").First(custStmtMsg, 1)

	if custStmtMsg == nil {
		return nil, fmt.Errorf("customer statement not found")
	}
	return custStmtMsg, nil
}
