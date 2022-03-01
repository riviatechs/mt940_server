package resolver

import (
	"context"

	"github.com/riviatechs/mt940_server/models"
)

type CabResolver struct{ *Resolver }

func (r *CabResolver) ID(ctx context.Context, obj *models.Cab) (int, error) {
	panic("not implemented")
}

func (r *CabResolver) Amount(ctx context.Context, obj *models.Cab) (float64, error) {
	panic("not implemented")
}
