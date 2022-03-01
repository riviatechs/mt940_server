package resolver

import (
	"context"

	"github.com/riviatechs/mt940_server/models"
)

type fwabResolver struct{ *Resolver }

func (r *fwabResolver) ID(ctx context.Context, obj *models.Fwab) (int, error) {
	panic("not implemented")
}

func (r *fwabResolver) Amount(ctx context.Context, obj *models.Fwab) (float64, error) {
	panic("not implemented")
}
