package resolver

import (
	"context"

	"github.com/riviatechs/mt940_server/graph/generated"
	"github.com/riviatechs/mt940_server/models"
	"github.com/riviatechs/mt940_server/util"
)

type FwabResolver struct{ *Resolver }

// Fwab returns generated.FwabResolver implementation.
func (r *Resolver) Fwab() generated.FwabResolver {
	return &FwabResolver{r}
}

func (r *FwabResolver) CustStmtMsgID(ctx context.Context, obj *models.Fwab) (int, error) {
	if obj == nil {
		return 0, nil
	}

	return int(obj.CustStmtMsgID), nil
}

func (r *FwabResolver) ID(ctx context.Context, obj *models.Fwab) (*int, error) {
	if obj == nil {
		return nil, nil
	}
	if obj.ID == nil {
		return nil, nil
	}
	id := int(*obj.ID)
	return &id, nil
}

func (r *FwabResolver) Amount(ctx context.Context, obj *models.Fwab) (float64, error) {
	if obj == nil {
		return 0, nil
	}

	return float64(obj.Amount), nil
}
func (r *FwabResolver) DateY(ctx context.Context, obj *models.Fwab) (string, error) {
	if obj == nil {
		return "", nil
	}

	return obj.DateY.Format(util.TimeFormat), nil
}
