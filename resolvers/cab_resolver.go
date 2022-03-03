package resolver

import (
	"context"

	"github.com/riviatechs/mt940_server/graph/generated"
	"github.com/riviatechs/mt940_server/models"
	"github.com/riviatechs/mt940_server/util"
)

type CabResolver struct{ *Resolver }

// Cab returns generated.CabResolver implementation.
func (r *Resolver) Cab() generated.CabResolver {
	return &CabResolver{r}
}

func (r *CabResolver) CustStmtMsgID(ctx context.Context, obj *models.Cab) (int, error) {
	if obj == nil {
		return 0, nil
	}

	return int(obj.CustStmtMsgID), nil

}

func (r *CabResolver) ID(ctx context.Context, obj *models.Cab) (*int, error) {
	if obj == nil {
		return nil, nil
	}
	if obj.ID == nil {
		return nil, nil
	}
	id := int(*obj.ID)
	return &id, nil
}

func (r *CabResolver) Amount(ctx context.Context, obj *models.Cab) (float64, error) {
	if obj == nil {
		return 0, nil
	}
	return float64(obj.Amount), nil
}

func (r *CabResolver) DateY(ctx context.Context, obj *models.Cab) (string, error) {
	if obj == nil {
		return "", nil
	}

	return obj.DateY.Format(util.TimeFormat), nil
}
