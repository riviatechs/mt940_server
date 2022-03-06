package resolver

import (
	"context"

	"github.com/riviatechs/mt940_server/db"
	"github.com/riviatechs/mt940_server/graph/generated"
	"github.com/riviatechs/mt940_server/models"
	"github.com/riviatechs/mt940_server/pdf"
)

type QueryResolver struct{ *Resolver }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &QueryResolver{r} }

func (r *Resolver) DownloadStatements(ctx context.Context, input *models.FilterInput) ([]*string, error) {
	s, err := pdf.GeneratePDF()
	if err != nil {
		return nil, err
	}
	return []*string{s}, nil
}

func (r *QueryResolver) Search(ctx context.Context, input string) ([]*models.ConfGroup, error) {
	return db.Search(ctx, input)
}

func (r *QueryResolver) StatementsFiltered(ctx context.Context, input *models.FilterInput) ([]*models.ConfGroup, error) {
	return db.StatementsFiltered(ctx, input)
}

func (r *QueryResolver) Statements(ctx context.Context) ([]*models.ConfGroup, error) {
	return db.Statements(ctx)
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
