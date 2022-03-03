package resolver

//go:generate go run github.com/99designs/gqlgen generate

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"github.com/riviatechs/mt940_server/graph/generated"
)

type Resolver struct{}

// Ai returns generated.AiResolver implementation.
func (r *Resolver) Ai() generated.AiResolver {
	return &AiResolver{r}
}

// Cab returns generated.CabResolver implementation.
func (r *Resolver) Cab() generated.CabResolver {
	return &CabResolver{r}
}

// Cb returns generated.CbResolver implementation.
func (r *Resolver) Cb() generated.CbResolver {
	return &CbResolver{r}
}

// CustStmtMsg returns generated.CustStmtMsgResolver implementation.
func (r *Resolver) CustStmtMsg() generated.CustStmtMsgResolver {
	return &CustStmtMsgResolver{r}
}

// Fwab returns generated.FwabResolver implementation.
func (r *Resolver) Fwab() generated.FwabResolver {
	return &FwabResolver{r}
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &MutationResolver{r} }

// Ob returns generated.ObResolver implementation.
func (r *Resolver) Ob() generated.ObResolver {
	return &ObResolver{r}
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &QueryResolver{r} }

// Sl returns generated.SlResolver implementation.
func (r *Resolver) Sl() generated.SlResolver { return &SlResolver{r} }

func (r *Resolver) CustStmtMsgInput() generated.CustStmtMsgInputResolver {
	return &CustStmtMsgInputResolver{r}
}

func (r *Resolver) SlInput() generated.SlInputResolver { return &SlInputResolver{r} }

func (r *Resolver) CabInput() generated.CabInputResolver { return &CabInputResolver{r} }

func (r *Resolver) CbInput() generated.CbInputResolver { return &CbInputResolver{r} }

func (r *Resolver) FwabInput() generated.FwabInputResolver { return &FwabInputResolver{r} }

func (r *Resolver) ObInput() generated.ObInputResolver { return &ObInputResolver{r} }

func (r *Resolver) AmountInput() generated.AmountInputResolver {
	return &AmountInputResolver{r}
}
