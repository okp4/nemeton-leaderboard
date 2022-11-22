package message

import (
	"okp4/nemeton-leaderboard/app/event"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PublishEventMessage struct {
	Event event.Event
}

type SubscribeEventMessage struct {
	From *primitive.ObjectID
}

type NewEventMessage struct {
	Event event.Event
}

type BrokenStreamMessage struct{}
