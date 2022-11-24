package nemeton

import "time"

type Task struct {
	ID          string    `bson:"id"`
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
