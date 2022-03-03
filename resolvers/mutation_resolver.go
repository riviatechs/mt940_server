package resolver

import (
	"context"

	"github.com/riviatechs/mt940_server/db"
	"github.com/riviatechs/mt940_server/graph/generated"
	"github.com/riviatechs/mt940_server/models"
)

type MutationResolver struct{ *Resolver }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &MutationResolver{r} }

func (r *MutationResolver) CreateCustStmtMsg(ctx context.Context, input models.CustStmtMsgInput) (*int, error) {
	return db.CreateCustStmtMsg(ctx, input)
}
