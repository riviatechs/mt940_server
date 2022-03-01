package resolver

import (
	"context"

	"github.com/riviatechs/mt940_server/models"
)

type CabInputResolver struct{ *Resolver }

func (r *CabInputResolver) Amount(ctx context.Context, obj *models.Cab, data float64) error {
	if obj == nil {
		return nil
	}

	obj.Amount = float32(data)
	return nil
}
