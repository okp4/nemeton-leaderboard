package nemeton

import "time"

const (
	taskTypeBasic        = "basic"
	taskTypeSubmission   = "submission"
	taskTypeUptime       = "uptime"
	taskTypeTweetNemeton = "tweet-nemeton"
)

type Task struct {
	ID          string    `bson:"id"`
	Type        string    `bson:"@type"`
	Name        string    `bson:"name"`
	Description string    `bson:"description"`
	StartDate   time.Time `bson:"startDate"`
	EndDate     time.Time `bson:"endDate"`
	Rewards     *int      `bson:"rewards,omitempty"`
}

func (t Task) Started() bool {
	return time.Now().After(t.StartDate)
}

func (t Task) Finished() bool {
	return time.Now().After(t.EndDate)
}

func (t Task) WithSubmission() bool {
	return t.Type == taskTypeSubmission
}
