package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"okp4/nemeton-leaderboard/app/nemeton"
	"okp4/nemeton-leaderboard/graphql/generated"
	"okp4/nemeton-leaderboard/graphql/model"
)

// Started is the resolver for the started field.
func (r *phaseResolver) Started(ctx context.Context, obj *nemeton.Phase) (bool, error) {
	panic(fmt.Errorf("not implemented: Started - started"))
}

// Finished is the resolver for the finished field.
func (r *phaseResolver) Finished(ctx context.Context, obj *nemeton.Phase) (bool, error) {
	panic(fmt.Errorf("not implemented: Finished - finished"))
}

// Blocks is the resolver for the blocks field.
func (r *phaseResolver) Blocks(ctx context.Context, obj *nemeton.Phase) (*model.BlockRange, error) {
	panic(fmt.Errorf("not implemented: Blocks - blocks"))
}

// Phase is the resolver for the phase field.
func (r *queryResolver) Phase(ctx context.Context, number int) (*nemeton.Phase, error) {
	return r.store.GetPhase(number), nil
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
func (r *queryResolver) Validator(ctx context.Context, cursor *string, rank *int, valoper *string, delegator *string, discord *string, twitter *string) (*model.Validator, error) {
	panic(fmt.Errorf("not implemented: Validator - validator"))
}

// Started is the resolver for the started field.
func (r *taskResolver) Started(ctx context.Context, obj *nemeton.Task) (bool, error) {
	panic(fmt.Errorf("not implemented: Started - started"))
}

// Finished is the resolver for the finished field.
func (r *taskResolver) Finished(ctx context.Context, obj *nemeton.Task) (bool, error) {
	panic(fmt.Errorf("not implemented: Finished - finished"))
}

// WithSubmission is the resolver for the withSubmission field.
func (r *taskResolver) WithSubmission(ctx context.Context, obj *nemeton.Task) (bool, error) {
	panic(fmt.Errorf("not implemented: WithSubmission - withSubmission"))
}

// Rewards is the resolver for the rewards field.
func (r *taskResolver) Rewards(ctx context.Context, obj *nemeton.Task) (*int, error) {
	panic(fmt.Errorf("not implemented: Rewards - rewards"))
}

// Phase returns generated.PhaseResolver implementation.
func (r *Resolver) Phase() generated.PhaseResolver { return &phaseResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Task returns generated.TaskResolver implementation.
func (r *Resolver) Task() generated.TaskResolver { return &taskResolver{r} }

type (
	phaseResolver struct{ *Resolver }
	queryResolver struct{ *Resolver }
	taskResolver  struct{ *Resolver }
)
