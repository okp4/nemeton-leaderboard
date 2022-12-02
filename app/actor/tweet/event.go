package tweet

import (
	"encoding/json"
	"time"
)

const NewTweetEventType = "new-tweet"

type NewTweetEvent struct {
	ID        string    `json:"id"`
	AuthorID  string    `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
	Text      string    `json:"text"`
	User      User      `json:"user"`
}

func (e *NewTweetEvent) Marshall() (map[string]interface{}, error) {
	var event map[string]interface{}
	data, err := json.Marshal(&e)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &event)
	return event, err
}

func Unmarshall(data map[string]interface{}) (*NewTweetEvent, error) {
	var event *NewTweetEvent
	d, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(d, &event)
	return event, err
}
