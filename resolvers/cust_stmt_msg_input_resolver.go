package resolver

import (
	"context"

	"github.com/riviatechs/mt940_server/models"
)

type CustStmtMsgInputResolver struct{ *Resolver }

func (r *CustStmtMsgInputResolver) Sl(ctx context.Context, obj *models.CustStmtMsg, data []*models.Sl) error {
	if obj == nil {
		return nil
	}
	obj.Sl = data
	return nil
}

func (r *CustStmtMsgInputResolver) Fwab(ctx context.Context, obj *models.CustStmtMsg, data []*models.Fwab) error {
	if obj == nil {
		return nil
	}

	obj.Fwab = data
	return nil
}
