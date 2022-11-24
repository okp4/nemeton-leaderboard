package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"okp4/nemeton-leaderboard/app/nemeton"
	"okp4/nemeton-leaderboard/graphql/generated"
	"okp4/nemeton-leaderboard/graphql/model"

	"github.com/cosmos/cosmos-sdk/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Picture is the resolver for the picture field.
func (r *identityResolver) Picture(ctx context.Context, obj *model.Identity) (*model.Link, error) {
	panic(fmt.Errorf("not implemented: Picture - picture"))
}

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
	return r.store.GetUnstartedPhases(), nil
}

// Finished is the resolver for the finished field.
func (r *phasesResolver) Finished(ctx context.Context, obj *model.Phases) ([]*nemeton.Phase, error) {
	return r.store.GetFinishedPhases(), nil
}

// Current is the resolver for the current field.
func (r *phasesResolver) Current(ctx context.Context, obj *model.Phases) (*nemeton.Phase, error) {
	return r.store.GetCurrentPhase(), nil
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
func (r *queryResolver) Board(ctx context.Context, search *string, first *int, after *primitive.ObjectID) (*model.BoardConnection, error) {
	panic(fmt.Errorf("not implemented: Board - board"))
}

// ValidatorCount is the resolver for the validatorCount field.
func (r *queryResolver) ValidatorCount(ctx context.Context) (int, error) {
	panic(fmt.Errorf("not implemented: ValidatorCount - validatorCount"))
}

// Validator is the resolver for the validator field.
func (r *queryResolver) Validator(ctx context.Context, cursor *primitive.ObjectID, rank *int, valoper types.ValAddress, delegator types.AccAddress, discord *string, twitter *string) (*nemeton.Validator, error) {
	if cursor != nil {
		return r.store.GetValidatorByID(ctx, *cursor)
	}
	if rank != nil {
		panic(fmt.Errorf("not implemented: Validator - validator"))
	}
	if !valoper.Empty() {
		return r.store.GetValidatorByValoper(ctx, valoper)
	}
	if !delegator.Empty() {
		return r.store.GetValidatorByDelegator(ctx, delegator)
	}
	if discord != nil {
		return r.store.GetValidatorByDiscord(ctx, *discord)
	}
	if twitter != nil {
		return r.store.GetValidatorByTwitter(ctx, *twitter)
	}
	return nil, fmt.Errorf("one option must be passed")
}

// Rank is the resolver for the rank field.
func (r *validatorResolver) Rank(ctx context.Context, obj *nemeton.Validator) (int, error) {
	panic(fmt.Errorf("not implemented: Rank - rank"))
}

// Identity is the resolver for the identity field.
func (r *validatorResolver) Identity(ctx context.Context, obj *nemeton.Validator) (*model.Identity, error) {
	panic(fmt.Errorf("not implemented: Identity - identity"))
}

// Status is the resolver for the status field.
func (r *validatorResolver) Status(ctx context.Context, obj *nemeton.Validator) (model.ValidatorStatus, error) {
	panic(fmt.Errorf("not implemented: Status - status"))
}

// Tasks is the resolver for the tasks field.
func (r *validatorResolver) Tasks(ctx context.Context, obj *nemeton.Validator) (*model.Tasks, error) {
	panic(fmt.Errorf("not implemented: Tasks - tasks"))
}

// MissedBlocks is the resolver for the missedBlocks field.
func (r *validatorResolver) MissedBlocks(ctx context.Context, obj *nemeton.Validator) ([]*model.BlockRange, error) {
	panic(fmt.Errorf("not implemented: MissedBlocks - missedBlocks"))
}

// Identity returns generated.IdentityResolver implementation.
func (r *Resolver) Identity() generated.IdentityResolver { return &identityResolver{r} }

// Phase returns generated.PhaseResolver implementation.
func (r *Resolver) Phase() generated.PhaseResolver { return &phaseResolver{r} }

// Phases returns generated.PhasesResolver implementation.
func (r *Resolver) Phases() generated.PhasesResolver { return &phasesResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Validator returns generated.ValidatorResolver implementation.
func (r *Resolver) Validator() generated.ValidatorResolver { return &validatorResolver{r} }

type (
	identityResolver  struct{ *Resolver }
	phaseResolver     struct{ *Resolver }
	phasesResolver    struct{ *Resolver }
	queryResolver     struct{ *Resolver }
	validatorResolver struct{ *Resolver }
)
