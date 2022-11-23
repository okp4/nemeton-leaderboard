package message

import (
	"okp4/nemeton-leaderboard/app/event"

	"github.com/asynkron/protoactor-go/actor"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PublishEventMessage struct {
	Event event.Event
}

type SubscribeEventMessage struct {
	PID  *actor.PID
	From *primitive.ObjectID
}

type NewEventMessage struct {
	Event event.Event
}

type BrokenStreamMessage struct{}
