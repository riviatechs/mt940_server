package resolver

import (
	"context"
	"time"

	"github.com/MalukiMuthusi/logger"
	"github.com/riviatechs/mt940_server/models"
	"github.com/riviatechs/mt940_server/util"
	"go.uber.org/zap"
)

type FwabInputResolver struct{ *Resolver }

func (r *FwabInputResolver) Amount(ctx context.Context, obj *models.Fwab, data float64) error {
	if obj == nil {
		return nil
	}
	obj.Amount = float32(data)
	return nil
}

func (r *FwabInputResolver) DateY(ctx context.Context, obj *models.Fwab, data string) error {
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
