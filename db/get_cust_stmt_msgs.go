package db

import (
	"context"

	"github.com/riviatechs/mt940_server/models"
)

func GetCustStmtMsgs(ctx context.Context) ([]*models.CustStmtMsg, error) {
	var custStmtMsgs []*models.CustStmtMsg

	Db.Preload("Ai").Preload("Ob").Preload("Sl").Preload("Cb").Preload("Cab").Preload("Fwab").Find(&custStmtMsgs)

	return custStmtMsgs, nil
}
