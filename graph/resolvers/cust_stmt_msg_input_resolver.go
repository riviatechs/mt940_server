package resolver

import (
	"context"

	"github.com/riviatechs/mt940_server/models"
)

type CustStmtMsgInputResolver struct{ *Resolver }

func (r *CustStmtMsgInputResolver) Sl(ctx context.Context, obj *models.CustStmtMsgInput, data []*models.SlInput) error {
	panic("not implemented")
}

func (r *CustStmtMsgInputResolver) Fwab(ctx context.Context, obj *models.CustStmtMsgInput, data []*models.TransInput) error {
	panic("not implemented")
}
