package resolver

import (
	"context"

	"github.com/riviatechs/mt940_server/models"
)

type SlInputResolver struct{ *Resolver }

func (r *SlInputResolver) Amount(ctx context.Context, obj *models.Sl, data float64) error {
	if obj == nil {
		return nil
	}
	obj.Amount = float32(data)
	return nil
}
