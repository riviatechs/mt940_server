package resolver

import (
	"context"

	"github.com/riviatechs/mt940_server/models"
)

type AmountInputResolver struct{ *Resolver }

func (r *AmountInputResolver) Lower(ctx context.Context, obj *models.Amount, data float64) error {
	if obj == nil {
		return nil
	}

	l := float32(data)
	obj.Lower = l

	return nil
}

func (r *AmountInputResolver) Upper(ctx context.Context, obj *models.Amount, data float64) error {
	if obj == nil {
		return nil
	}

	l := float32(data)
	obj.Upper = l

	return nil
}
