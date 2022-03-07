package resolver

import (
	"context"

	"github.com/riviatechs/mt940_server/graph/generated"
	"github.com/riviatechs/mt940_server/models"
)

type DownloadInputResolver struct {
	*Resolver
}

func (r *Resolver) DownloadInput() generated.DownloadInputResolver {
	return &DownloadInputResolver{r}
}

func (r *DownloadInputResolver) FilterInput(ctx context.Context, obj *models.DownloadInput, data *models.FilterInput) error {
	if obj == nil {
		return nil
	}

	obj.Filters = data
	return nil
}

func (r *DownloadInputResolver) FieldsInput(ctx context.Context, obj *models.DownloadInput, data *models.FieldsInput) error {
	if obj == nil {
		return nil
	}

	obj.Fields = data
	return nil
}
