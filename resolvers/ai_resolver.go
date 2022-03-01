package resolver

import (
	"context"

	"github.com/riviatechs/mt940_server/models"
)

type aiResolver struct{ *Resolver }

func (r *aiResolver) ID(ctx context.Context, obj *models.Ai) (int, error) {
	panic("not implemented")
}

func (r *aiResolver) CreatedAt(ctx context.Context, obj *models.Ai) (*string, error) {
	panic("not implemented")
}

func (r *aiResolver) UpdatedAt(ctx context.Context, obj *models.Ai) (*string, error) {
	panic("not implemented")
}
