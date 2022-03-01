package resolver

import (
	"context"

	"github.com/riviatechs/mt940_server/models"
)

type FwabResolver struct{ *Resolver }

func (r *FwabResolver) ID(ctx context.Context, obj *models.Fwab) (*int, error) {
	if obj == nil {
		return nil, nil
	}
	if obj.ID == nil {

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
