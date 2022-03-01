package resolver

import (
	"context"

	"github.com/riviatechs/mt940_server/models"
)

type AiResolver struct{ *Resolver }

func (r *AiResolver) ID(ctx context.Context, obj *models.Ai) (int, error) {
	panic("not implemented")
}

func (r *AiResolver) CreatedAt(ctx context.Context, obj *models.Ai) (*string, error) {
	panic("not implemented")
}

func (r *AiResolver) UpdatedAt(ctx context.Context, obj *models.Ai) (*string, error) {
	panic("not implemented")
}
