package resolver

import (
	"context"
	"time"

	"github.com/riviatechs/mt940_server/graph/generated"
	"github.com/riviatechs/mt940_server/models"
	"github.com/riviatechs/mt940_server/util"
)

type SlInputResolver struct{ *Resolver }

func (r *Resolver) SlInput() generated.SlInputResolver { return &SlInputResolver{r} }

func (r *SlInputResolver) CustStmtMsgID(ctx context.Context, obj *models.SlInput, data int) error {
	if obj == nil {
		return nil
	}

	obj.CustStmtMsgID = uint(data)
	return nil
}

func (r *SlInputResolver) Amount(ctx context.Context, obj *models.SlInput, data float64) error {
	if obj == nil {
		return nil
	}
	obj.Amount = float32(data)
	return nil
}

func (r *SlInputResolver) ValueDate(ctx context.Context, obj *models.SlInput, data *string) error {
	if obj == nil {
		return nil
	}

	if data == nil {
		return nil
	}

	valueDate := util.FormatDate(*data)
	t, err := time.Parse(util.TimeFormat, valueDate)
	if err != nil {
		tt := time.Now()
		obj.ValueDate = tt
		return nil
	}

	obj.ValueDate = t
	return nil
}

func (r *SlInputResolver) EntryDate(ctx context.Context, obj *models.SlInput, data *string) error {
	if obj == nil {
		return nil
	}

	if data == nil {
		return nil
	}

	entryDate := util.FormatDate(*data)
	t, err := time.Parse(util.TimeFormat, entryDate)
	if err != nil {
		tt := time.Now()
		obj.EntryDate = &tt
		return nil
	}

	obj.EntryDate = &t
	return nil
}
