package resolver

import (
	"context"

	"github.com/riviatechs/mt940_server/models"
)

type obResolver struct{ *Resolver }

func (r *obResolver) ID(ctx context.Context, obj *models.Ob) (int, error) {
	panic("not implemented")
}
