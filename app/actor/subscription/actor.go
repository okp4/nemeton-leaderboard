package subscription

import (
	"context"
	"fmt"
	"time"

	"okp4/nemeton-leaderboard/app/util"

	"okp4/nemeton-leaderboard/graphql"

	"okp4/nemeton-leaderboard/app/actor/synchronization"
	"okp4/nemeton-leaderboard/app/actor/tweet"
	"okp4/nemeton-leaderboard/app/event"
	"okp4/nemeton-leaderboard/app/message"
	"okp4/nemeton-leaderboard/app/nemeton"
	"okp4/nemeton-leaderboard/app/offset"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const ownerOffset = "subscription"

type Actor struct {
	store       *nemeton.Store
	ctx         context.Context
	eventPID    *actor.PID
	offsetStore *offset.Store
	lastHeight  *int64
}

func NewSubscriber(mongoURI, dbName string, eventPID *actor.PID) (*Actor, error) {
	ctx := context.Background()
	store, err := nemeton.NewStore(ctx, mongoURI, dbName)
	if err != nil {
		return nil, err
	}

	offsetStore, err := offset.NewStore(ctx, mongoURI, dbName, ownerOffset)
	if err != nil {
		return nil, err
	}

	return &Actor{
		store:       store,
		ctx:         ctx,
		eventPID:    eventPID,
		offsetStore: offsetStore,
	}, nil
}

func (a *Actor) Receive(ctx actor.Context) {
	switch e := ctx.Message().(type) {
	case *actor.Started:
		var from *primitive.ObjectID
		value, _ := a.offsetStore.Get(a.ctx)
		switch v := value.(type) {
		case primitive.ObjectID:
			from = &v
		default:
			from = nil
		}

		log.Info().Msg("🕵️ Start looking for new event")
		ctx.Send(a.eventPID, &message.SubscribeEventMessage{
			PID:  ctx.Self(),
			From: from,
		})
	case *message.NewEventMessage:
		a.receiveNewEvent(e.Event)
	case *actor.Restarting, *actor.Stopping:
		log.Info().Msg("✋ Stop looking new event")
		ctx.Send(a.eventPID, &message.UnsubscribeEventMessage{
			PID: ctx.Self(),
		})
		if err := a.offsetStore.Close(context.Background()); err != nil {
			log.Err(err).Msg("❌ Couldn't properly close offset store")
		}
		if err := a.store.Close(context.Background()); err != nil {
			log.Err(err).Msg("❌ Couldn't properly close offset store")
		}
	}
}

func (a *Actor) receiveNewEvent(e event.Event) {
	log.Info().Str("type", e.Type).Msg("📦 Receive event")
	switch e.Type {
	case synchronization.NewBlockEventType:
		a.handleNewBlockEvent(e.Data)
	case graphql.GenTXSubmittedEventType:
		a.handleGenTXSubmittedEvent(e.Date, e.Data)
	case graphql.ValidatorRegisteredEventType:
		a.handleValidatorRegisteredEvent(e.Data)
	case graphql.ValidatorUpdatedEventType:
		a.handleValidatorUpdatedEvent(e.Data)
	case graphql.ValidatorRemovedEventType:
		a.handleValidatorRemovedEvent(e.Data)
	case tweet.NewTweetEventType:
		a.handleNewTweetEvent(e.Date, e.Data)
	case graphql.TaskSubmittedEventType:
		a.handleTaskSubmittedEvent(e.Data)
	case graphql.TaskCompletedEventType:
		a.handleTaskCompletedEvent(e.Data)
	case graphql.RegisterURLEventType:
		a.handleRegisterURLEvent(e.Date, e.Data)
	default:
		log.Warn().Msg("⚠️ No event handler for this event.")
	}
	if err := a.offsetStore.Save(a.ctx, e.ID); err != nil {
		log.Panic().Err(err).Msg("❌ Failed save offset state.")
	}
}

func (a *Actor) handleNewBlockEvent(data map[string]interface{}) {
	e := &synchronization.NewBlockEvent{}
	if err := e.Unmarshal(data); err != nil {
		log.Panic().Err(err).Msg("❌ Failed unmarshal event to NewBlockEvent")
		return
	}

	logger := log.With().Time("blockTime", e.Time).Int64("height", e.Height).Logger()
	logger.Info().Msg("Handle NewBlock event")
	a.lastHeight = &e.Height

	consensusAddr := make([]types.ConsAddress, len(e.Signatures))
	for i, signature := range e.Signatures {
		consensusAddr[i] = signature.GetValidatorAddress()
	}

	if err := a.store.UpdateValidatorUptime(a.ctx, consensusAddr, e.Height); err != nil {
		logger.Panic().Err(err).Msg("🤕 Failed update validator uptime.")
	}

	if err := a.store.CompleteNodeSetupTask(a.ctx, e.Time, consensusAddr); err != nil {
		logger.Panic().Err(err).Msg("🤕 Failed update validator node setup task.")
	}

	if err := a.store.CompleteVoteProposalTask(a.ctx, e.Time, e.MsgVotes); err != nil {
		logger.Panic().Err(err).Msg("🤕 Failed complete vote proposal task.")
	}

	if err := a.store.UpdatePhaseBlocks(a.ctx, e.Time, e.Height); err != nil {
		logger.Panic().Err(err).Msg("🤕 Failed update phase block range.")
	}

	previousPhase, err := a.store.GetPreviousPhaseByBlock(a.ctx, e.Height)
	if err != nil {
		log.Panic().Err(err).Msg("🤕 Failed get previous phase.")
	}

	phase := a.store.GetCurrentPhaseAt(e.Time)
	blockRange, err := a.store.GetPhaseBlocks(a.ctx, phase.Number)
	if err != nil {
		log.Panic().Err(err).Msg("🤕 Could not request block range.")
	}

	if previousPhase != nil && previousPhase.Number < phase.Number {
		log.Info().Int("oldPhase", previousPhase.Number).Msg("⏱️ It's the previous phase ended")
		a.handlePhaseEnded(previousPhase)
	}

	if blockRange != nil && blockRange.To-blockRange.From == 1 {
		log.Info().Int("newPhase", phase.Number).Msg("⏱️ It's a new phase started! ")
		a.handlePhaseStarted(phase)
	}
}

func (a *Actor) handleGenTXSubmittedEvent(when time.Time, data map[string]interface{}) {
	log.Info().Interface("event", data).Msg("Handle GenTXSubmitted event")

	e := &graphql.GenTXSubmittedEvent{}
	if err := e.Unmarshal(data); err != nil {
		log.Panic().Err(err).Msg("❌ Failed unmarshal event to GenTXSubmitted")
		return
	}

	msgCreateVal, err := util.ParseGenTX(e.GenTX)
	if err != nil {
		log.Panic().Err(err).Msg("❌ Failed unmarshal gentx")
	}

	if err := a.store.CreateGentxValidator(context.Background(), when, msgCreateVal, e.Discord, e.Country, e.Twitter); err != nil {
		log.Err(err).Interface("data", data).Msg("🤕 Couldn't create validator")
	}
}

func (a *Actor) handleValidatorRegisteredEvent(data map[string]interface{}) {
	log.Info().Interface("event", data).Msg("Handle ValidatorRegistered event")

	e := &graphql.ValidatorRegisteredEvent{}
	if err := e.Unmarshal(data); err != nil {
		log.Panic().Err(err).Msg("❌ Failed unmarshal event to ValidatorRegistered")
		return
	}

	if a.lastHeight == nil {
		log.Err(fmt.Errorf("doesn't have a last height")).
			Interface("data", data).
			Msg("🤕 Couldn't register validator")
	}

	if err := a.store.RegisterValidator(
		context.Background(),
		e.Valoper,
		e.Delegator,
		e.Valcons,
		e.Description,
		e.Discord,
		e.Country,
		e.Twitter,
		*a.lastHeight,
	); err != nil {
		log.Err(err).Interface("data", data).Msg("🤕 Couldn't register validator")
	}
}

func (a *Actor) handleValidatorUpdatedEvent(data map[string]interface{}) {
	log.Info().Interface("event", data).Msg("Handle ValidatorUpdated event")

	e := &graphql.ValidatorUpdatedEvent{}
	if err := e.Unmarshal(data); err != nil {
		log.Panic().Err(err).Msg("❌ Failed unmarshal event to ValidatorUpdated")
		return
	}

	if err := a.store.UpdateValidator(
		context.Background(),
		e.Delegator,
		e.Valoper,
		e.Valcons,
		e.Description,
		e.Discord,
		e.Country,
		e.Twitter,
	); err != nil {
		log.Err(err).Interface("data", data).Msg("🤕 Couldn't update validator")
	}
}

func (a *Actor) handleValidatorRemovedEvent(data map[string]interface{}) {
	log.Info().Interface("event", data).Msg("Handle ValidatorRemoved event")

	e := &graphql.ValidatorRemovedEvent{}
	if err := e.Unmarshal(data); err != nil {
		log.Panic().Err(err).Msg("❌ Failed unmarshal event to ValidatorRemoved")
		return
	}

	if err := a.store.RemoveValidator(a.ctx, e.Validator); err != nil {
		log.Err(err).Interface("data", data).Msg("🤕 Couldn't remove validator")
	}
}

func (a *Actor) handleNewTweetEvent(when time.Time, data map[string]interface{}) {
	log.Info().Interface("event", data).Msg("Handle NewTweet event")

	e := &tweet.NewTweetEvent{}
	if err := e.Unmarshal(data); err != nil {
		log.Panic().Err(err).Msg("❌ Failed unmarshal event to NewTweetEvent")
		return
	}
	phase := a.store.GetCurrentPhaseAt(when)
	for _, task := range phase.Tasks {
		if task.Type == nemeton.TaskTypeTweetNemeton && task.InProgressAt(when) {
			if !e.CreatedAt.After(task.StartDate) || !e.CreatedAt.Before(task.EndDate) {
				log.Warn().Time("startDate", task.StartDate).
					Time("endDate", task.EndDate).
					Time("tweetDate", e.CreatedAt).
					Msg("🐦 Tweet has been posted before or after the task time.")
				continue
			}

			err := a.store.CompleteTweetTask(a.ctx, when, e.User.Username, phase, task)
			if err != nil {
				log.Panic().Err(err).Msg("❌ Could not complete tweet task")
			}
			return // We consider that there is only one tweet task by phase
		}
	}
}

func (a *Actor) handlePhaseEnded(phase *nemeton.Phase) {
	err := a.store.CompleteValidatorsUptimeForPhase(a.ctx, phase)
	if err != nil {
		log.Panic().Err(err).Msg("❌⏱️ An error occurs fetch uptime validators.")
		return
	}
	log.Info().Int("phaseNumber", phase.Number).Msg("✅ Uptime points for phase has been set.")
}

func (a *Actor) handlePhaseStarted(phase *nemeton.Phase) {
	// TODO: handle phase started
}

func (a *Actor) handleRegisterURLEvent(when time.Time, data map[string]interface{}) {
	log.Info().Interface("event", data).Msg("Handle RegisterURL event")

	e := &graphql.RegisterURLEvent{}
	if err := e.Unmarshal(data); err != nil {
		log.Panic().Err(err).Msg("❌ Failed unmarshal event to RegisterURLEvent")
		return
	}

	if err := a.store.RegisterValidatorURL(a.ctx, when, e.Type, e.Validator, e.URL, e.Points); err != nil {
		log.Err(err).Interface("data", data).Msg("🤕 Couldn't register/update validator url")
	}
}

func (a *Actor) handleTaskSubmittedEvent(data map[string]interface{}) {
	log.Info().Interface("event", data).Msg("Handle TaskSubmitted event")

	e := &graphql.TaskSubmittedEvent{}
	if err := e.Unmarshal(data); err != nil {
		log.Panic().Err(err).Msg("❌ Failed unmarshal event to TaskSubmittedEvent")
		return
	}

	if err := a.store.ManualSubmitTask(a.ctx, e.Validator, e.Phase, e.Task); err != nil {
		log.Err(err).Interface("data", data).Msg("🤕 Couldn't manually submit task")
	}
}

func (a *Actor) handleTaskCompletedEvent(data map[string]interface{}) {
	log.Info().Interface("event", data).Msg("Handle TaskCompleted event")

	e := &graphql.TaskCompletedEvent{}
	if err := e.Unmarshal(data); err != nil {
		log.Panic().Err(err).Msg("❌ Failed unmarshal event to TaskCompletedEvent")
		return
	}

	if err := a.store.ManualCompleteTask(a.ctx, e.Validator, e.Phase, e.Task, e.Points); err != nil {
		log.Err(err).Interface("data", data).Msg("🤕 Couldn't manually complete task")
	}
}
