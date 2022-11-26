package graphql

import (
	"okp4/nemeton-leaderboard/app/keybase"
	"okp4/nemeton-leaderboard/app/nemeton"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	store         *nemeton.Store
	keybaseClient *keybase.Client
}

func NewResolver(store *nemeton.Store, keybaseClient *keybase.Client) *Resolver {
	return &Resolver{
		store:         store,
		keybaseClient: keybaseClient,
	}
}
