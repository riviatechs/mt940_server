package resolver

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"

	"github.com/riviatechs/mt940_server/graph/generated"
	"github.com/riviatechs/mt940_server/models"
)

type Resolver struct{}

func (r *aiResolver) ID(ctx context.Context, obj *models.Ai) (int, error) {
	panic("not implemented")
}

func (r *cabResolver) ID(ctx context.Context, obj *models.Cab) (int, error) {
	panic("not implemented")
}

func (r *cabResolver) Amount(ctx context.Context, obj *models.Cab) (float64, error) {
	panic("not implemented")
}

func (r *cbResolver) ID(ctx context.Context, obj *models.Cb) (int, error) {
	panic("not implemented")
}

func (r *cbResolver) Amount(ctx context.Context, obj *models.Cb) (float64, error) {
	panic("not implemented")
}

func (r *custStmtMsgResolver) ID(ctx context.Context, obj *models.CustStmtMsg) (*int, error) {
	panic("not implemented")
}

func (r *custStmtMsgResolver) Sl(ctx context.Context, obj *models.CustStmtMsg) ([]*models.Sl, error) {
	panic("not implemented")
}

func (r *custStmtMsgResolver) Fwab(ctx context.Context, obj *models.CustStmtMsg) ([]*models.Fwab, error) {
	panic("not implemented")
}

func (r *fwabResolver) ID(ctx context.Context, obj *models.Fwab) (int, error) {
	panic("not implemented")
}

func (r *fwabResolver) Amount(ctx context.Context, obj *models.Fwab) (float64, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateCustStmtMsg(ctx context.Context, input *models.CustStmtMsgInput) (*int, error) {
	panic("not implemented")
}

func (r *obResolver) ID(ctx context.Context, obj *models.Ob) (int, error) {
	panic("not implemented")
}

func (r *obResolver) Amount(ctx context.Context, obj *models.Ob) (float64, error) {
	panic("not implemented")
}

func (r *queryResolver) CustStmtMsg(ctx context.Context, id int) (*models.CustStmtMsg, error) {
	panic("not implemented")
}

func (r *slResolver) ID(ctx context.Context, obj *models.Sl) (*int, error) {
	panic("not implemented")
}

func (r *slResolver) Amount(ctx context.Context, obj *models.Sl) (float64, error) {
	panic("not implemented")
}

// Ai returns generated.AiResolver implementation.
func (r *Resolver) Ai() generated.AiResolver { return &aiResolver{r} }

// Cab returns generated.CabResolver implementation.
func (r *Resolver) Cab() generated.CabResolver { return &cabResolver{r} }

// Cb returns generated.CbResolver implementation.
func (r *Resolver) Cb() generated.CbResolver { return &cbResolver{r} }

// CustStmtMsg returns generated.CustStmtMsgResolver implementation.
func (r *Resolver) CustStmtMsg() generated.CustStmtMsgResolver { return &custStmtMsgResolver{r} }

// Fwab returns generated.FwabResolver implementation.
func (r *Resolver) Fwab() generated.FwabResolver { return &fwabResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Ob returns generated.ObResolver implementation.
func (r *Resolver) Ob() generated.ObResolver { return &obResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Sl returns generated.SlResolver implementation.
func (r *Resolver) Sl() generated.SlResolver { return &slResolver{r} }

type aiResolver struct{ *Resolver }
type cabResolver struct{ *Resolver }
type cbResolver struct{ *Resolver }
type custStmtMsgResolver struct{ *Resolver }
type fwabResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type obResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type slResolver struct{ *Resolver }
