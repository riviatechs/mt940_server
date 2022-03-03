package resolver

import (
	"context"

	"github.com/riviatechs/mt940_server/graph/generated"
	"github.com/riviatechs/mt940_server/models"
)

type CustStmtMsgInputResolver struct{ *Resolver }

func (r *Resolver) CustStmtMsgInput() generated.CustStmtMsgInputResolver {
	return &CustStmtMsgInputResolver{r}
}

func (r *CustStmtMsgInputResolver) Sl(ctx context.Context, obj *models.CustStmtMsgInput, data []*models.SlInput) error {
	if obj == nil {
		return nil
	}
	obj.Sl = data
	return nil
}

func (r *CustStmtMsgInputResolver) Fwab(ctx context.Context, obj *models.CustStmtMsgInput, data []*models.FwabInput) error {
	if obj == nil {
		return nil
	}

	obj.Fwab = data
	return nil
}
