package tweet

import (
	"encoding/json"
	"time"
)

const NewTweetEventType = "new-tweet"

type NewTweetEvent struct {
	ID       string    `json:"id"`
	AuthorID string    `json:"author_id"`
	Time     time.Time `json:"time"`
	Text     string    `json:"text"`
	User     User      `json:"user"`
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
