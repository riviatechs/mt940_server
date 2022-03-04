package resolver

import (
	"context"

	"github.com/riviatechs/mt940_server/graph/generated"
	"github.com/riviatechs/mt940_server/models"
	"github.com/riviatechs/mt940_server/util"
)

type ConfirmationResolver struct{ *Resolver }

func (r *Resolver) Confirmation() generated.ConfirmationResolver {
	return &ConfirmationResolver{r}
}

func (r *ConfirmationResolver) Amount(ctx context.Context, obj *models.Confirmation) (float64, error) {
	if obj == nil {
		return 0, nil
	}
	return float64(obj.Amount), nil
}

func (r *ConfirmationResolver) DateTime(ctx context.Context, obj *models.Confirmation) (string, error) {
	if obj == nil {
		return "", nil
	}

	return obj.DateTime.Format(util.TimeFormat), nil
}

func (r *ConfirmationResolver) ID(ctx context.Context, obj *models.Confirmation) (int, error) {
	if obj == nil {
		return 0, nil
	}

	return int(obj.ID), nil
}
