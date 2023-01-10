package nemeton

import "time"

const (
	taskTypeGentx        = "gentx"
	taskTypeSubmission   = "submission"
	taskTypeNodeSetup    = "node-setup"
	taskTypeUptime       = "uptime"
	TaskTypeTweetNemeton = "tweet-nemeton"
	TaskTypeRPC          = "rpc"
	TaskTypeSnapshots    = "snapshots"
	TaskTypeDashboard    = "dashboard"

	taskParamMaxPoints = "max-points"
)

type TaskState struct {
	Completed    bool   `bson:"completed"`
	EarnedPoints uint64 `bson:"points"`
	Submitted    bool   `bson:"submitted"`
}

type Task struct {
	ID          string                 `bson:"id"`
	Type        string                 `bson:"@type"`
	Name        string                 `bson:"name"`
	Description string                 `bson:"description"`
	StartDate   time.Time              `bson:"startDate"`
	EndDate     time.Time              `bson:"endDate"`
	Rewards     *uint64                `bson:"rewards,omitempty"`
	Params      map[string]interface{} `bson:"params,omitempty"`
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

func (t Task) StartedAt(at time.Time) bool {
	return at.After(t.StartDate)
}

func (t Task) FinishedAt(at time.Time) bool {
	return at.After(t.EndDate)
}

func (t Task) InProgressAt(at time.Time) bool {
	return t.StartedAt(at) && !t.FinishedAt(at)
}

func (t Task) WithSubmission() bool {
	return t.Type == taskTypeSubmission ||
		t.Type == TaskTypeRPC ||
		t.Type == TaskTypeSnapshots ||
		t.Type == TaskTypeDashboard
}

func (t Task) GetParamMaxPoints() *uint64 {
	if v, ok := t.Params[taskParamMaxPoints]; ok {
		if maxPoints, ok := v.(int64); ok {
			p := uint64(maxPoints)
			return &p
		}
	}
	return nil
}

func makeTask(
	ttype, id, name, description string,
	start, end time.Time,
	rewards *uint64,
	params map[string]interface{},
) Task {
	return Task{
		ID:          id,
		Type:        ttype,
		Name:        name,
		Description: description,
		StartDate:   start,
		EndDate:     end,
		Rewards:     rewards,
		Params:      params,
	}
}

func makeGentxTask(id, name, description string, start, end time.Time, rewards uint64) Task {
	return makeTask(taskTypeGentx, id, name, description, start, end, &rewards, nil)
}

func makeSubmissionTask(id, name, description string, start, end time.Time) Task {
	return makeTask(taskTypeSubmission, id, name, description, start, end, nil, nil)
}

func makeNodeSetupTask(id, name, description string, start, end time.Time, rewards uint64) Task {
	return makeTask(taskTypeNodeSetup, id, name, description, start, end, &rewards, nil)
}

func makeUptimeTask(id, name, description string, start, end time.Time, maxPoints uint64) Task {
	return makeTask(taskTypeUptime, id, name, description, start, end, nil, map[string]interface{}{
		taskParamMaxPoints: maxPoints,
	})
}

func makeTweetNemetonTask(id, name, description string, start, end time.Time, rewards uint64) Task {
	return makeTask(TaskTypeTweetNemeton, id, name, description, start, end, &rewards, nil)
}

func makeRPCTask(id, name, description string, start, end time.Time, rewards uint64) Task {
	return makeTask(TaskTypeRPC, id, name, description, start, end, &rewards, nil)
}

func makeSnapshotsTask(id, name, description string, start, end time.Time, rewards uint64) Task {
	return makeTask(TaskTypeSnapshots, id, name, description, start, end, &rewards, nil)
}

func makeDashboardTask(id, name, description string, start, end time.Time, maxPoints uint64) Task {
	return makeTask(TaskTypeDashboard, id, name, description, start, end, nil, map[string]interface{}{
		taskParamMaxPoints: maxPoints,
	})
}
