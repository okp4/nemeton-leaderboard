package nemeton

import "time"

const (
	taskTypeBasic        = "basic"
	taskTypeGentx        = "gentx"
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
	Rewards     *uint64   `bson:"rewards,omitempty"`
}

func (t Task) Started() bool {
	return time.Now().After(t.StartDate)
}

func (t Task) Finished() bool {
	return time.Now().After(t.EndDate)
}

func (t Task) InProgress() bool {
	return t.Started() && !t.Finished()
}

func (t Task) WithSubmission() bool {
	return t.Type == taskTypeSubmission
}

type TaskState struct {
	Completed    bool   `bson:"completed"`
	EarnedPoints uint64 `bson:"points"`
}
