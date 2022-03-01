package resolver

import (
	"context"

	"github.com/riviatechs/mt940_server/models"
)

type cabResolver struct{ *Resolver }

func (r *cabResolver) ID(ctx context.Context, obj *models.Cab) (int, error) {
	panic("not implemented")
}

func (r *cabResolver) Amount(ctx context.Context, obj *models.Cab) (float64, error) {
	panic("not implemented")
}


