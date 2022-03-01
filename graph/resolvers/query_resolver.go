package resolver

import (
	"context"

	"github.com/riviatechs/mt940_server/models"
)

type queryResolver struct{ *Resolver }

func (r *queryResolver) CustStmtMsg(ctx context.Context, id int) (*models.CustStmtMsg, error) {
	panic("not implemented")
}
