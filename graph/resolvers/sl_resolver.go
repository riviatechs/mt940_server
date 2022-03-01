package resolver

import (
	"context"

	"github.com/riviatechs/mt940_server/models"
)

type slResolver struct{ *Resolver }

func (r *slResolver) ID(ctx context.Context, obj *models.Sl) (*int, error) {
	panic("not implemented")
}

func (r *slResolver) Amount(ctx context.Context, obj *models.Sl) (float64, error) {
	panic("not implemented")
}
