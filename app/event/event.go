package event

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	id      primitive.ObjectID `bson:"_id"`
	evtType string             `bson:"@type"`
	date    time.Time
	data    map[string]interface{}
}

func NewEvent(evtType string, data map[string]interface{}) Event {
	return Event{
		evtType: evtType,
		date:    time.Now(),
		data:    data,
	}
}

func (e Event) ID() string {
	return e.id.String()
}

func (e Event) Type() string {
	return e.evtType
}
