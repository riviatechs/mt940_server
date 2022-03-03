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

type CbInputResolver struct{ *Resolver }

func (r *Resolver) CbInput() generated.CbInputResolver { return &CbInputResolver{r} }

func (r *CbInputResolver) CustStmtMsgID(ctx context.Context, obj *models.CbInput, data int) error {
	if obj == nil {
		return nil
	}

	obj.CustStmtMsgID = uint(data)
	return nil
}

func (r *CbInputResolver) Amount(ctx context.Context, obj *models.CbInput, data float64) error {
	if obj == nil {
		return nil
	}
	obj.Amount = float32(data)
	return nil
}

func (r *CbInputResolver) DateY(ctx context.Context, obj *models.CbInput, data string) error {
	if obj == nil {
		return nil
	}

	t, err := time.Parse(util.TimeFormat, data)
	if err != nil {
		logger.Info("Date", zap.Error(err))
		return fmt.Errorf("invalidate date")
	}

	obj.DateY = t
	return nil
}
