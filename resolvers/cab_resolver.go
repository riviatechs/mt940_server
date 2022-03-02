package resolver

import (
	"context"

	"github.com/riviatechs/mt940_server/models"
)

type CabResolver struct{ *Resolver }

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