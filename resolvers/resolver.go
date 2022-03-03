package resolver

//go:generate go run github.com/99designs/gqlgen generate

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"github.com/riviatechs/mt940_server/graph/generated"
)

type Resolver struct{}

func (r *Resolver) AmountInput() generated.AmountInputResolver {
	return &AmountInputResolver{r}
}

func (r *Resolver) AiInput() generated.AiInputResolver {
	return &AiInputResolver{r}
}
