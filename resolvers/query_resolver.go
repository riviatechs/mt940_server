package resolver

import (
	"context"

	"github.com/riviatechs/mt940_server/db"
	"github.com/riviatechs/mt940_server/models"
)

type queryResolver struct{ *Resolver }

func (r *queryResolver) CustStmtMsg(ctx context.Context, id int) (*models.CustStmtMsg, error) {
	return db.GetCustStmtMsg(ctx, id)
}

func (r *queryResolver) CustStmtMsgs(ctx context.Context) ([]*models.CustStmtMsg, error) {
	return db.GetCustStmtMsgs(ctx)
}
