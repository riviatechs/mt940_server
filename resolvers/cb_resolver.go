package resolver

import (
	"context"

	"github.com/riviatechs/mt940_server/models"
)

type CbResolver struct{ *Resolver }

func (r *CbResolver) ID(ctx context.Context, obj *models.Cb) (*int, error) {
	if obj == nil {
		return nil, nil
	}
	if obj.ID == nil {
		return nil, nil
	}
	id := int(*obj.ID)
	return &id, nil
}

func (r *CbResolver) Amount(ctx context.Context, obj *models.Cb) (float64, error) {
	if obj == nil {
		return 0, nil
	}
	return float64(obj.Amount), nil
}
