package resolver

import (
	"context"

	"github.com/riviatechs/mt940_server/models"
)

type custStmtMsgResolver struct{ *Resolver }

func (r *custStmtMsgResolver) ID(ctx context.Context, obj *models.CustStmtMsg) (*int, error) {
	panic("not implemented")
}

func (r *custStmtMsgResolver) Sl(ctx context.Context, obj *models.CustStmtMsg) ([]*models.Sl, error) {
	panic("not implemented")
}

func (r *custStmtMsgResolver) Fwab(ctx context.Context, obj *models.CustStmtMsg) ([]*models.Fwab, error) {
	panic("not implemented")
}

func (r *custStmtMsgResolver) Ai(ctx context.Context, obj *models.CustStmtMsg) (*models.Ai, error) {
	panic("not implemented")
}

func (r custStmtMsgResolver) Ob(ctx context.Context, obj *models.CustStmtMsg) (*models.Ob, error) {
	panic("not implemented")
}

func (r custStmtMsgResolver) Cab(ctx context.Context, obj *models.CustStmtMsg) (*models.Cab, error) {
	panic("not implemented")
}
