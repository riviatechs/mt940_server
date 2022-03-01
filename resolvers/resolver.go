package resolver

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"github.com/riviatechs/mt940_server/graph/generated"
	"github.com/riviatechs/mt940_server/models"
)

type Resolver struct{}

// Ai returns generated.AiResolver implementation.
func (r *Resolver) Ai() generated.AiResolver {
	return &aiResolver{r}
}

// Cab returns generated.CabResolver implementation.
func (r *Resolver) Cab() generated.CabResolver {
	return &cabResolver{r}
}

// Cb returns generated.CbResolver implementation.
func (r *Resolver) Cb() generated.CbResolver {
	return &cbResolver{r}
}

// CustStmtMsg returns generated.CustStmtMsgResolver implementation.
func (r *Resolver) CustStmtMsg() generated.CustStmtMsgResolver {
	return &custStmtMsgResolver{r}
}

// Fwab returns generated.FwabResolver implementation.
func (r *Resolver) Fwab() {
	panic("not implemented")
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Ob returns generated.ObResolver implementation.
func (r *Resolver) Ob() (*models.Ob, error) {
	panic("not implemented")
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Sl returns generated.SlResolver implementation.
func (r *Resolver) Sl() generated.SlResolver { return &slResolver{r} }

func (r *Resolver) SlInput() generated.SlInputResolver { return &SlInputResolver{r} }

func (r *Resolver) CustStmtMsgInput() generated.SlResolver { return &slResolver{r} }
