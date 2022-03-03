package resolver

import (
	"context"

	"github.com/riviatechs/mt940_server/graph/generated"
	"github.com/riviatechs/mt940_server/models"
	"github.com/riviatechs/mt940_server/util"
)

type SlResolver struct{ *Resolver }

// Sl returns generated.SlResolver implementation.
func (r *Resolver) Sl() generated.SlResolver { return &SlResolver{r} }

func (r *SlResolver) CustStmtMsgID(ctx context.Context, obj *models.Sl) (int, error) {
	if obj == nil {
		return 0, nil
	}

	return int(obj.CustStmtMsgID), nil
}

func (r *SlResolver) ID(ctx context.Context, obj *models.Sl) (*int, error) {
	if obj == nil {
		return nil, nil
	}
	if obj.ID == nil {
		return nil, nil
	}
	id := int(*obj.ID)
	return &id, nil
}

func (r *SlResolver) Amount(ctx context.Context, obj *models.Sl) (float64, error) {
	if obj == nil {
		return 0, nil
	}
	return float64(obj.Amount), nil
}

func (r *SlResolver) EntryDate(ctx context.Context, obj *models.Sl) (*string, error) {
	if obj == nil {
		return nil, nil
	}
	if obj.EntryDate == nil {
		return nil, nil
	}

	entryDate := obj.EntryDate.Format(util.TimeFormat)

	return &entryDate, nil
}

func (r *SlResolver) ValueDate(ctx context.Context, obj *models.Sl) (*string, error) {
	if obj == nil {
		return nil, nil
	}
	entryDate := obj.ValueDate.Format(util.TimeFormat)
	return &entryDate, nil
}
