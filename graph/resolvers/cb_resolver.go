package resolver

import (
	"context"

	"github.com/riviatechs/mt940_server/models"
)

type CbResolver struct{ *Resolver }

func (r *CbResolver) ID(ctx context.Context, obj *models.Cb) (int, error) {
	panic("not implemented")
}

func (r *CbResolver) Amount(ctx context.Context, obj *models.Cb) (float64, error) {
	panic("not implemented")
}
