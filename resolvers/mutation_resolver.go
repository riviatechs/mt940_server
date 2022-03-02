package resolver

import (
	"context"

	"github.com/riviatechs/mt940_server/db"
	"github.com/riviatechs/mt940_server/models"
)

type MutationResolver struct{ *Resolver }

func (r *MutationResolver) CreateCustStmtMsg(ctx context.Context, input models.CustStmtMsg) (*int, error) {
	return db.CreateCustStmtMsg(ctx, input)
}
