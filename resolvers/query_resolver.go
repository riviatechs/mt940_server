package resolver

import (
	"context"

	"github.com/riviatechs/mt940_server/db"
	"github.com/riviatechs/mt940_server/graph/generated"
	"github.com/riviatechs/mt940_server/models"
)

type QueryResolver struct{ *Resolver }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &QueryResolver{r} }

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

func (r *QueryResolver) GetStmtLineGroupedByDate(ctx context.Context) ([]*models.SlGroups, error) {
	return db.GetStmtLineGroupedByDate(ctx)
}

func (r *QueryResolver) GetStmtLinesFilterByAmount(ctx context.Context, amount models.Amount) ([]*models.SlGroups, error) {
	return db.GetStmtLinesFilterByPeriod(ctx, amount)
}
