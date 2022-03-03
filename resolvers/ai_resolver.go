package resolver

import (
	"context"

	"github.com/riviatechs/mt940_server/graph/generated"
	"github.com/riviatechs/mt940_server/models"
	"github.com/riviatechs/mt940_server/util"
)

type AiResolver struct{ *Resolver }

// Ai returns generated.AiResolver implementation.
func (r *Resolver) Ai() generated.AiResolver {
	return &AiResolver{r}
}

func (r *AiResolver) CustStmtMsgID(ctx context.Context, obj *models.Ai) (int, error) {
	if obj == nil {
		return 0, nil
	}

	return int(obj.CustStmtMsgID), nil
}

func (r *AiResolver) ID(ctx context.Context, obj *models.Ai) (*int, error) {
	if obj == nil {
		return nil, nil
	}
	if obj.ID == nil {
		return nil, nil
	}
	id := int(*obj.ID)
	return &id, nil
}

func (r *AiResolver) CreatedAt(ctx context.Context, obj *models.Ai) (*string, error) {
	if obj == nil {
		return nil, nil
	}

	if obj.CreatedAt == nil {
		return nil, nil
	}

	t := obj.CreatedAt
	tStr := t.Format(util.TimeFormat)
	return &tStr, nil
}

func (r *AiResolver) UpdatedAt(ctx context.Context, obj *models.Ai) (*string, error) {
	if obj == nil {
		return nil, nil
	}
	if obj.UpdatedAt == nil {
		return nil, nil
	}
	t := obj.UpdatedAt
	tStr := t.Format(util.TimeFormat)
	return &tStr, nil
}

type AiInputResolver struct{ *Resolver }

func (r *AiInputResolver) CustStmtMsgID(ctx context.Context, obj *models.AiInput, data int) error {
	if obj == nil {
		return nil
	}

	obj.CustStmtMsgID = uint(data)
	return nil
}
