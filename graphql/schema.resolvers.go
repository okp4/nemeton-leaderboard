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

// Blocks is the resolver for the blocks field.
func (r *phaseResolver) Blocks(ctx context.Context, obj *nemeton.Phase) (*model.BlockRange, error) {
	panic(fmt.Errorf("not implemented: Blocks - blocks"))
}

// All is the resolver for the all field.
func (r *phasesResolver) All(ctx context.Context, obj *model.Phases) ([]*nemeton.Phase, error) {
	return r.store.GetAllPhases(), nil
}

// Ongoing is the resolver for the ongoing field.
func (r *phasesResolver) Ongoing(ctx context.Context, obj *model.Phases) ([]*nemeton.Phase, error) {
	panic(fmt.Errorf("not implemented: Ongoing - ongoing"))
}

// Finished is the resolver for the finished field.
func (r *phasesResolver) Finished(ctx context.Context, obj *model.Phases) ([]*nemeton.Phase, error) {
	panic(fmt.Errorf("not implemented: Finished - finished"))
}

// Current is the resolver for the current field.
func (r *phasesResolver) Current(ctx context.Context, obj *model.Phases) (*nemeton.Phase, error) {
	panic(fmt.Errorf("not implemented: Current - current"))
}

// Phase is the resolver for the phase field.
func (r *queryResolver) Phase(ctx context.Context, number int) (*nemeton.Phase, error) {
	return r.store.GetPhase(number), nil
}

// Phases is the resolver for the phases field.
func (r *queryResolver) Phases(ctx context.Context) (*model.Phases, error) {
	return &model.Phases{}, nil
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

// Phase returns generated.PhaseResolver implementation.
func (r *Resolver) Phase() generated.PhaseResolver { return &phaseResolver{r} }

// Phases returns generated.PhasesResolver implementation.
func (r *Resolver) Phases() generated.PhasesResolver { return &phasesResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type (
	phaseResolver  struct{ *Resolver }
	phasesResolver struct{ *Resolver }
	queryResolver  struct{ *Resolver }
)
