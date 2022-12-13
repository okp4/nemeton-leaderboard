package subscription

import (
	"context"
	"time"

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
	case tweet.NewTweetEventType:
		a.handleNewTweetEvent(e.Date, e.Data)
	default:
		log.Warn().Msg("⚠️ No event handler for this event.")
	}
	if err := a.offsetStore.Save(a.ctx, e.ID); err != nil {
		log.Panic().Err(err).Msg("❌ Failed save offset state.")
	}
}

func (a *Actor) handleNewBlockEvent(data map[string]interface{}) {
	log.Info().Interface("event", data).Msg("Handle NewBlock event")

	e, err := synchronization.Unmarshall(data)
	if err != nil {
		log.Panic().Err(err).Msg("❌ Failed unmarshall event to NewBlockEvent")
		return
	}

	consensusAddr := make([]types.ConsAddress, len(e.Signatures))
	for i, signature := range e.Signatures {
		consensusAddr[i] = signature.GetValidatorAddress()
	}

	if err := a.store.UpdateValidatorUptime(a.ctx, consensusAddr, e.Height); err != nil {
		log.Panic().Err(err).Msg("🤕 Failed update validator uptime.")
	}
}

func (a *Actor) handleGenTXSubmittedEvent(when time.Time, data map[string]interface{}) {
	log.Info().Interface("event", data).Msg("Handle GenTXSubmitted event")

	e, err := graphql.Unmarshall(data)
	if err != nil {
		log.Panic().Err(err).Msg("❌ Failed unmarshall event to GenTXSubmitted")
		return
	}

	if err := a.store.CreateValidator(context.Background(), when, e.Discord, e.Country, e.Twitter, e.GenTX); err != nil {
		log.Err(err).Interface("data", data).Msg("🤕 Couldn't create validator")
	}
}

func (a *Actor) handleNewTweetEvent(when time.Time, data map[string]interface{}) {
	log.Info().Interface("event", data).Msg("Handle NewTweet event")

	e, err := tweet.Unmarshall(data)
	if err != nil {
		log.Panic().Err(err).Msg("❌ Failed unmarshall event to NewTweetEvent")
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
