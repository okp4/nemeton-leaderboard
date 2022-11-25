package nemeton

import "time"

type Phase struct {
	Number      int       `bson:"_id"`
	Name        string    `bson:"name"`
	Description string    `bson:"description"`
	StartDate   time.Time `bson:"startDate"`
	EndDate     time.Time `bson:"endDate"`
	Tasks       []Task    `bson:"tasks"`
}

func (p Phase) Started() bool {
	return time.Now().After(p.StartDate)
}

func (p Phase) Finished() bool {
	return time.Now().After(p.EndDate)
}

func (p Phase) InProgress() bool {
	return p.Started() && !p.Finished()
}
