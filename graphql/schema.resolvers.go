package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"okp4/nemeton-leaderboard/graphql/generated"
	"okp4/nemeton-leaderboard/graphql/model"
)

// Phase is the resolver for the phase field.
func (r *queryResolver) Phase(ctx context.Context, number int) (*model.Phase, error) {
	panic(fmt.Errorf("not implemented: Phase - phase"))
}

// Phases is the resolver for the phases field.
func (r *queryResolver) Phases(ctx context.Context) (*model.Phases, error) {
	panic(fmt.Errorf("not implemented: Phases - phases"))
}

// Board is the resolver for the board field.
func (r *queryResolver) Board(ctx context.Context, search *string, first *int, after *string) (*model.BoardConnection, error) {
	panic(fmt.Errorf("not implemented: Board - board"))
}

// ValidatorCount is the resolver for the validatorCount field.
func (r *queryResolver) ValidatorCount(ctx context.Context) (int, error) {
	panic(fmt.Errorf("not implemented: ValidatorCount - validatorCount"))
}

// Validator is the resolver for the validator field.
// nolint: lll
func (r *queryResolver) Validator(ctx context.Context, cursor *string, rank *int, valoper *string, delegator *string, discord *string, twitter *string) (*model.Validator, error) {
	panic(fmt.Errorf("not implemented: Validator - validator"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
