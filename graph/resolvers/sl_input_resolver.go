package resolver

import (
	"context"

	"github.com/riviatechs/mt940_server/models"
)

type SlInputResolver struct{ *Resolver }

func (r *SlInputResolver) Amount(ctx context.Context, obj *models.SlInput, data float64) error {
	panic("not implemented")
}
