package resolver

import (
	"context"

	"github.com/riviatechs/mt940_server/models"
)

type FwabResolver struct{ *Resolver }

func (r *FwabResolver) ID(ctx context.Context, obj *models.Fwab) (int, error) {
	panic("not implemented")
}

func (r *FwabResolver) Amount(ctx context.Context, obj *models.Fwab) (float64, error) {
	panic("not implemented")
}
