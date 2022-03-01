package resolver

import (
	"context"

	"github.com/riviatechs/mt940_server/models"
)

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateCustStmtMsg(ctx context.Context, input models.CustStmtMsgInput) (*int, error) {
	panic("not implemented")
}
