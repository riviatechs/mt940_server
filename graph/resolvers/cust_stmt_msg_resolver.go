package resolver

import (
	"context"

	"github.com/riviatechs/mt940_server/models"
)

type CustStmtMsgResolver struct{ *Resolver }

func (r *CustStmtMsgResolver) ID(ctx context.Context, obj *models.CustStmtMsg) (*int, error) {
	panic("not implemented")
}

func (r *CustStmtMsgResolver) Sl(ctx context.Context, obj *models.CustStmtMsg) ([]*models.Sl, error) {
	panic("not implemented")
}

func (r *CustStmtMsgResolver) Fwab(ctx context.Context, obj *models.CustStmtMsg) ([]*models.Fwab, error) {
	panic("not implemented")
}

func (r *CustStmtMsgResolver) Ai(ctx context.Context, obj *models.CustStmtMsg) (*models.Ai, error) {
	panic("not implemented")
}

func (r CustStmtMsgResolver) Ob(ctx context.Context, obj *models.CustStmtMsg) (*models.Ob, error) {
	panic("not implemented")
}

func (r CustStmtMsgResolver) Cab(ctx context.Context, obj *models.CustStmtMsg) (*models.Cab, error) {
	panic("not implemented")
}
