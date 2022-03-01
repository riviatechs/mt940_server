package resolver

import (
	"context"

	"github.com/riviatechs/mt940_server/models"
)

type cbResolver struct{ *Resolver }

func (r *cbResolver) ID(ctx context.Context, obj *models.Cb) (int, error) {
	panic("not implemented")
}

func (r *cbResolver) Amount(ctx context.Context, obj *models.Cb) (float64, error) {
	panic("not implemented")
}
