package resolver

import (
	"context"

	"github.com/riviatechs/mt940_server/models"
)

type slResolver struct{ *Resolver }

func (r *slResolver) ID(ctx context.Context, obj *models.Sl) (*int, error) {
	if obj == nil {
		return nil, nil
	}
	if obj.ID == nil {
		return nil, nil
	}
	id := int(*obj.ID)
	return &id, nil
}

func (r *slResolver) Amount(ctx context.Context, obj *models.Sl) (float64, error) {
	if obj == nil {
		return 0, nil
	}
	return float64(obj.Amount), nil
}
