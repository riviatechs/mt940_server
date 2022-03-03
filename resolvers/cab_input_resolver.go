package resolver

import (
	"context"
	"fmt"
	"time"

	"github.com/MalukiMuthusi/logger"
	"github.com/riviatechs/mt940_server/graph/generated"
	"github.com/riviatechs/mt940_server/models"
	"github.com/riviatechs/mt940_server/util"
	"go.uber.org/zap"
)

type CabInputResolver struct{ *Resolver }

func (r *Resolver) CabInput() generated.CabInputResolver { return &CabInputResolver{r} }

func (r *CabInputResolver) CustStmtMsgID(ctx context.Context, obj *models.CabInput, data int) error {
	if obj == nil {
		return nil
	}

	obj.CustStmtMsgID = uint(data)
	return nil
}

func (r *CabInputResolver) Amount(ctx context.Context, obj *models.CabInput, data float64) error {
	if obj == nil {
		return nil
	}

	obj.Amount = float32(data)
	return nil
}

func (r *CabInputResolver) DateY(ctx context.Context, obj *models.CabInput, data string) error {
	if obj == nil {
		return nil
	}
	dateY, err := time.Parse(util.TimeFormat, data)
	if err != nil {
		logger.Info("Date", zap.Error(err))
		return fmt.Errorf("invalid date")
	}

	obj.DateY = dateY
	return nil
}
