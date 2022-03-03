package resolver

import (
	"context"
	"fmt"
	"time"

	"github.com/MalukiMuthusi/logger"
	"github.com/riviatechs/mt940_server/models"
	"github.com/riviatechs/mt940_server/util"
	"go.uber.org/zap"
)

type ObResolver struct{ *Resolver }

func (r *ObResolver) ID(ctx context.Context, obj *models.Ob) (*int, error) {
	if obj == nil {
		return nil, nil
	}
	if obj.ID == nil {
		return nil, nil
	}
	id := int(*obj.ID)
	return &id, nil
}

func (r *ObResolver) Amount(ctx context.Context, obj *models.Ob) (float64, error) {
	if obj == nil {
		return 0, nil
	}
	return float64(obj.Amount), nil
}

type ObInputResolver struct{ *Resolver }

func (r *ObInputResolver) Amount(ctx context.Context, obj *models.Ob, data float64) error {
	if obj == nil {
		return nil
	}
	obj.Amount = float32(data)
	return nil
}

func (r *ObResolver) DateY(ctx context.Context, obj *models.Ob) (string, error) {
	if obj == nil {
		return "", nil
	}

	return obj.DateY.Format(util.TimeFormat), nil
}

func (r *ObInputResolver) DateY(ctx context.Context, obj *models.Ob, data string) error {
	if obj == nil {
		return nil
	}

	t, err := time.Parse(util.TimeFormat, data)
	if err != nil {
		logger.Info("Date", zap.Error(err))
		return fmt.Errorf("invalid date format")
	}
	obj.DateY = t
	return nil
}
