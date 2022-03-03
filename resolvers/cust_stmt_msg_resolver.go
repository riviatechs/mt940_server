package resolver

import (
	"context"

	"github.com/riviatechs/mt940_server/graph/generated"
	"github.com/riviatechs/mt940_server/models"
)

type CustStmtMsgResolver struct{ *Resolver }

// CustStmtMsg returns generated.CustStmtMsgResolver implementation.
func (r *Resolver) CustStmtMsg() generated.CustStmtMsgResolver {
	return &CustStmtMsgResolver{r}
}

func (r *CustStmtMsgResolver) ID(ctx context.Context, obj *models.CustStmtMsg) (*int, error) {
	if obj == nil {
		return nil, nil
	}
	if obj.ID == nil {
		return nil, nil
	}
	id := int(*obj.ID)
	return &id, nil
}

func (r *CustStmtMsgResolver) Sl(ctx context.Context, obj *models.CustStmtMsg) ([]*models.Sl, error) {
	if obj == nil {
		return nil, nil
	}

	return obj.Sl, nil
}

func (r *CustStmtMsgResolver) Fwab(ctx context.Context, obj *models.CustStmtMsg) ([]*models.Fwab, error) {
	if obj == nil {
		return nil, nil
	}
	return obj.Fwab, nil
}

func (r *CustStmtMsgResolver) Ai(ctx context.Context, obj *models.CustStmtMsg) (*models.Ai, error) {
	if obj == nil {
		return nil, nil
	}

	return &obj.Ai, nil
}

func (r CustStmtMsgResolver) Ob(ctx context.Context, obj *models.CustStmtMsg) (*models.Ob, error) {
	if obj == nil {
		return nil, nil
	}

	return &obj.Ob, nil
}

func (r CustStmtMsgResolver) Cab(ctx context.Context, obj *models.CustStmtMsg) (*models.Cab, error) {
	if obj == nil {
		return nil, nil
	}
	return obj.Cab, nil
}
