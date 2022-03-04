package resolver

import (
	"context"

	"github.com/riviatechs/mt940_server/graph/generated"
	"github.com/riviatechs/mt940_server/models"
	"github.com/riviatechs/mt940_server/util"
)

type ConfResolver struct {
	*Resolver
}

func (r *Resolver) ConfGroup() generated.ConfGroupResolver {
	return &ConfResolver{r}
}

func (r *ConfResolver) DateTime(ctx context.Context, obj *models.ConfGroup) (*string, error) {
	if obj == nil {
		return nil, nil
	}

	t := obj.DateTime.Format(util.TimeFormat)
	return &t, nil
}
