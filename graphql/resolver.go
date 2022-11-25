package graphql

import (
	"okp4/nemeton-leaderboard/app/nemeton"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	store *nemeton.Store
}

func NewResolver(store *nemeton.Store) *Resolver {
	return &Resolver{
		store: store,
	}
}
