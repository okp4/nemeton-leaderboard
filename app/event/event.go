package event

import "time"

type Event struct {
	id      string `bson:"_id"`
	evtType string `bson:"@type"`
	date    time.Time
	data    interface{}
}

func NewEvent(evtType string, data interface{}) Event {
	return Event{
		evtType: evtType,
		date:    time.Now(),
		data:    data,
	}
}
