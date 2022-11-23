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
