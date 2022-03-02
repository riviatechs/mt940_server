package db

import (
	"context"

	"github.com/riviatechs/mt940_server/models"
)

func GetCustStmtMsgByTrn(ctx context.Context, trn string) (*models.CustStmtMsg, error) {
	var custStmtMsg models.CustStmtMsg
	Db.Preload("Ai").Preload("Ob").Preload("Sl").Preload("Cb").Preload("Cab").Preload("Fwab").Where("trn = ?", trn).First(&custStmtMsg)

	return &custStmtMsg, nil
}
