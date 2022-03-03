package resolver

import (
	"context"
	"time"

	"github.com/MalukiMuthusi/logger"
	"github.com/riviatechs/mt940_server/graph/generated"
	"github.com/riviatechs/mt940_server/models"
	"github.com/riviatechs/mt940_server/util"
	"go.uber.org/zap"
)

type FwabInputResolver struct{ *Resolver }

func (r *Resolver) FwabInput() generated.FwabInputResolver { return &FwabInputResolver{r} }

func (r *FwabInputResolver) CustStmtMsgID(ctx context.Context, obj *models.FwabInput, data int) error {
	if obj == nil {
		return nil
	}

	obj.CustStmtMsgID = uint(data)
	return nil
}

func (r *FwabInputResolver) Amount(ctx context.Context, obj *models.FwabInput, data float64) error {
	if obj == nil {
		return nil
	}
	obj.Amount = float32(data)
	return nil
}

func (r *FwabInputResolver) DateY(ctx context.Context, obj *models.FwabInput, data string) error {
	if obj == nil {
		return nil
	}

	t, err := time.Parse(util.TimeFormat, data)
	if err != nil {
		logger.Info("Date", zap.Error(err))
		return nil
	}

	obj.DateY = t
	return nil
}
