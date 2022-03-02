package resolver

import (
	"context"

	"github.com/riviatechs/mt940_server/db"
	"github.com/riviatechs/mt940_server/models"
)

type QueryResolver struct{ *Resolver }

func (r *QueryResolver) CustStmtMsg(ctx context.Context, id int) (*models.CustStmtMsg, error) {
	return db.GetCustStmtMsg(ctx, id)
}

func (r *QueryResolver) CustStmtMsgs(ctx context.Context) ([]*models.CustStmtMsg, error) {
	return db.GetCustStmtMsgs(ctx)
}

func (r *QueryResolver) GetStatementLines(ctx context.Context) ([]*models.Sl, error) {
	return db.GetStatementLines(ctx)
}

func (r *QueryResolver) GetCustStmtMsgByTrn(ctx context.Context, trn string) (*models.CustStmtMsg, error) {
	return db.GetCustStmtMsgByTrn(ctx, trn)
}
