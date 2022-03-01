package resolver

import (
	"context"

	"github.com/riviatechs/mt940_server/models"
)

type FwabInputResolver struct{ *Resolver }

func (r *FwabInputResolver) Amount(ctx context.Context, obj *models.Fwab, data float64) error {
	if obj == nil {
		return nil
	}
	obj.Amount = float32(data)
	return nil
}
