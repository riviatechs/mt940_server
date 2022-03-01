package resolver

import (
	"context"

	"github.com/riviatechs/mt940_server/models"
)

type CbInputResolver struct{ *Resolver }

func (r *CbInputResolver) Amount(ctx context.Context, obj *models.Cb, data float64) error {
	if obj == nil {
		return nil
	}
	obj.Amount = float32(data)
	return nil
}
