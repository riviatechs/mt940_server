package resolver

import (
	"context"

	"github.com/riviatechs/mt940_server/models"
)

type ObResolver struct{ *Resolver }

func (r *ObResolver) ID(ctx context.Context, obj *models.Ob) (*int, error) {
	if obj == nil {
		return nil, nil
	}
	if obj.ID == nil {
		return nil, nil
	}
	id := int(*obj.ID)
	return &id, nil
}

func (r *ObResolver) Amount(ctx context.Context, obj *models.Ob) (float64, error) {
	if obj == nil {
		return 0, nil
	}
	return float64(obj.Amount), nil
}

type ObInputResolver struct{ *Resolver }

func (r *ObInputResolver) Amount(ctx context.Context, obj *models.Ob, data float64) error {
	if obj == nil {
		return nil
	}
	obj.Amount = float32(data)
	return nil
}
