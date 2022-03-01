package resolver

import (
	"context"

	"github.com/riviatechs/mt940_server/models"
)

type ObResolver struct{ *Resolver }

func (r *ObResolver) ID(ctx context.Context, obj *models.Ob) (int, error) {
	panic("not implemented")
}

func (r *ObResolver) Amount(ctx context.Context, obj *models.Ob) (float64, error) {
	panic("not implemented")
}
