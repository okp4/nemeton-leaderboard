package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"okp4/nemeton-leaderboard/app/event"
	"okp4/nemeton-leaderboard/app/message"
	"okp4/nemeton-leaderboard/app/nemeton"
	"okp4/nemeton-leaderboard/graphql/generated"
	"okp4/nemeton-leaderboard/graphql/model"

	"github.com/cosmos/cosmos-sdk/types"
)

// Picture is the resolver for the picture field.
func (r *identityResolver) Picture(ctx context.Context, obj *model.Identity) (*model.Link, error) {
	picture, err := r.keybaseClient.LookupPicture(ctx, obj.Kid)
	if err != nil {
		return nil, err
	}

	var link *model.Link
	if picture != nil {
		link = &model.Link{Href: picture}
	}
	return link, nil
}

// SubmitValidatorGenTx is the resolver for the submitValidatorGenTX field.
func (r *mutationResolver) SubmitValidatorGenTx(ctx context.Context, twitter *string, discord string, country string, gentx map[string]interface{}) (*string, error) {
	evt := GenTXSubmittedEvent{
		Twitter: twitter,
		Discord: discord,
		Country: country,
		GenTX:   gentx,
	}
	rawEvt, err := evt.Marshall()
	if err != nil {
		return nil, err
	}

	r.actorCTX.Send(
		r.eventStore,
		&message.PublishEventMessage{
			Event: event.NewEvent(
				GenTXSubmittedEventType,
				rawEvt,
			),
		},
	)
	return nil, nil
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
func (r *queryResolver) Board(ctx context.Context, search *string, first *int, after *nemeton.Cursor) (*model.BoardConnection, error) {
	validators, hasNext, err := r.store.GetBoard(ctx, search, *first, after)
	if err != nil {
		return nil, err
	}

	edges := make([]*model.ValidatorEdge, 0, len(validators))
	for _, validator := range validators {
		edges = append(edges, &model.ValidatorEdge{
			Cursor: validator.Cursor(),
			Node:   validator,
		})
	}

	var startCursor *nemeton.Cursor
	var endCursor *nemeton.Cursor
	if len(edges) > 0 {
		startCursor = edges[0].Cursor
		endCursor = edges[len(validators)-1].Cursor
	}

	return &model.BoardConnection{
		Edges: edges,
		PageInfo: &model.PageInfo{
			StartCursor: startCursor,
			EndCursor:   endCursor,
			HasNextPage: hasNext,
			Count:       len(validators),
		},
	}, nil
}

// ValidatorCount is the resolver for the validatorCount field.
func (r *queryResolver) ValidatorCount(ctx context.Context) (int, error) {
	count, err := r.store.CountValidators(ctx)
	return int(count), err
}

// Validator is the resolver for the validator field.
func (r *queryResolver) Validator(ctx context.Context, cursor *nemeton.Cursor, rank *int, valoper types.ValAddress, delegator types.AccAddress, discord *string, twitter *string) (*nemeton.Validator, error) {
	if cursor != nil {
		return r.store.GetValidatorByCursor(ctx, *cursor)
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

// ForPhase is the resolver for the forPhase field.
func (r *tasksResolver) ForPhase(ctx context.Context, obj *model.Tasks, number int) (*model.PerPhaseTasks, error) {
	panic(fmt.Errorf("not implemented: ForPhase - forPhase"))
}

// Rank is the resolver for the rank field.
func (r *validatorResolver) Rank(ctx context.Context, obj *nemeton.Validator) (int, error) {
	return r.store.GetValidatorRank(ctx, *obj.Cursor())
}

// Identity is the resolver for the identity field.
func (r *validatorResolver) Identity(ctx context.Context, obj *nemeton.Validator) (*model.Identity, error) {
	if obj.Identity == nil {
		return nil, nil
	}

	return &model.Identity{
		Kid: *obj.Identity,
	}, nil
}

// Status is the resolver for the status field.
func (r *validatorResolver) Status(ctx context.Context, obj *nemeton.Validator) (model.ValidatorStatus, error) {
	panic(fmt.Errorf("not implemented: Status - status"))
}

// Tasks is the resolver for the tasks field.
func (r *validatorResolver) Tasks(ctx context.Context, obj *nemeton.Validator) (*model.Tasks, error) {
	result := &model.Tasks{
		CompletedCount: 0,
		StartedCount:   0,
		FinishedCount:  0,
	}
	for _, phase := range append(r.store.GetFinishedPhases(), r.store.GetCurrentPhase()) {
		perPhase := &model.PerPhaseTasks{
			CompletedCount: 0,
			StartedCount:   0,
			FinishedCount:  0,
			Phase:          phase,
		}
		for i, task := range phase.Tasks {
			mappedState := &model.BasicTaskState{
				Task: &phase.Tasks[i],
			}
			if state := obj.Task(phase.Number, task.ID); state != nil {
				mappedState.Completed = state.Completed
				mappedState.EarnedPoints = state.EarnedPoints
			}

			if mappedState.Completed {
				perPhase.CompletedCount++
			}
			if task.Started() {
				perPhase.StartedCount++
			}
			if task.Finished() {
				perPhase.FinishedCount++
			}
			perPhase.Tasks = append(perPhase.Tasks, mappedState)
		}

		result.CompletedCount += perPhase.CompletedCount
		result.StartedCount += perPhase.StartedCount
		result.FinishedCount += perPhase.FinishedCount
		result.PerPhase = append(result.PerPhase, perPhase)
	}

	return result, nil
}

// MissedBlocks is the resolver for the missedBlocks field.
func (r *validatorResolver) MissedBlocks(ctx context.Context, obj *nemeton.Validator) ([]*model.BlockRange, error) {
	panic(fmt.Errorf("not implemented: MissedBlocks - missedBlocks"))
}

// Identity returns generated.IdentityResolver implementation.
func (r *Resolver) Identity() generated.IdentityResolver { return &identityResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Phase returns generated.PhaseResolver implementation.
func (r *Resolver) Phase() generated.PhaseResolver { return &phaseResolver{r} }

// Phases returns generated.PhasesResolver implementation.
func (r *Resolver) Phases() generated.PhasesResolver { return &phasesResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Tasks returns generated.TasksResolver implementation.
func (r *Resolver) Tasks() generated.TasksResolver { return &tasksResolver{r} }

// Validator returns generated.ValidatorResolver implementation.
func (r *Resolver) Validator() generated.ValidatorResolver { return &validatorResolver{r} }

type (
	identityResolver  struct{ *Resolver }
	mutationResolver  struct{ *Resolver }
	phaseResolver     struct{ *Resolver }
	phasesResolver    struct{ *Resolver }
	queryResolver     struct{ *Resolver }
	tasksResolver     struct{ *Resolver }
	validatorResolver struct{ *Resolver }
)
