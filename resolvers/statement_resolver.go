package resolver

import (
	"context"

	"github.com/riviatechs/mt940_server/graph/generated"
	"github.com/riviatechs/mt940_server/models"
)

type StatementResolver struct{ *Resolver }

func (r *Resolver) Statement() generated.StatementResolver {
	return &StatementResolver{r}
}

func (r *StatementResolver) ID(ctx context.Context, obj *models.Statement) (int, error) {
	if obj == nil {
		return 0, nil
	}
	return int(obj.ID), nil
}

func (r *StatementResolver) Ob(ctx context.Context, obj *models.Statement) (float64, error) {
	if obj == nil {
		return 0, nil
	}
	b := obj.Ob
	return float64(b), nil
}

func (r *StatementResolver) Cb(ctx context.Context, obj *models.Statement) (float64, error) {
	if obj == nil {
		return 0, nil
	}
	b := obj.Cb
	return float64(b), nil
}
