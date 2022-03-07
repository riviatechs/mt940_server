package resolver

import (
	"context"

	"github.com/riviatechs/mt940_server/db"
	"github.com/riviatechs/mt940_server/download"
	"github.com/riviatechs/mt940_server/graph/generated"
	"github.com/riviatechs/mt940_server/models"
)

type QueryResolver struct{ *Resolver }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &QueryResolver{r} }

func (r *QueryResolver) Download(ctx context.Context, input models.DownloadInput) (*string, error) {
	return download.HandleDownload(ctx, input)
}

func (r *QueryResolver) Search(ctx context.Context, input string) ([]*models.ConfGroup, error) {
	return db.Search(ctx, input)
}

func (r *QueryResolver) StatementsFiltered(ctx context.Context, input *models.FilterInput) ([]*models.ConfGroup, error) {
	stmts, err := db.StatementsFiltered(ctx, input)
	if err != nil {
		return nil, err
	}

	confs, err := db.GroupStmtsByDate(stmts)
	if err != nil {
		return nil, err
	}

	return confs, nil
}

func (r *QueryResolver) Statements(ctx context.Context) ([]*models.ConfGroup, error) {
	stmts, err := db.Statements(ctx)
	if err != nil {
		return nil, err
	}

	confs, err := db.GroupStmtsByDate(stmts)
	if err != nil {
		return nil, err
	}

	return confs, nil
}

func (r *Resolver) GetStmtLinesFilterByDc(ctx context.Context, dcInput models.DCInput) ([]*models.SlGroups, error) {
	return db.FilterByDC(ctx, dcInput)
}

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
