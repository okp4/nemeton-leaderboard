package event

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Type string             `bson:"@type"`
	Date time.Time
	Data map[string]interface{}
}

func NewEvent(evtType string, data map[string]interface{}) Event {
	return Event{
		Type: evtType,
		Date: time.Now(),
		Data: data,
	}
}
