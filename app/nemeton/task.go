package nemeton

import "time"

const (
	taskTypeGentx        = "gentx"
	taskTypeSubmission   = "submission"
	taskTypeNodeSetup    = "node-setup"
	taskTypeUptime       = "uptime"
	TaskTypeTweetNemeton = "tweet-nemeton"
	taskTypeRPC          = "rpc"
	taskTypeSnapshots    = "snapshots"
	taskTypeDashboard    = "dashboard"

	taskParamUptimeMaxPoints = "max-points"
)

type TaskState struct {
	Completed    bool   `bson:"completed"`
	EarnedPoints uint64 `bson:"points"`
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

func (t Task) GetUptimeMaxPoints() *uint64 {
	if v, ok := t.Params[taskParamUptimeMaxPoints]; ok {
		if maxPoints, ok := v.(uint64); ok {
			return &maxPoints
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
		taskParamUptimeMaxPoints: maxPoints,
	})
}

func makeTweetNemetonTask(id, name, description string, start, end time.Time, rewards uint64) Task {
	return makeTask(TaskTypeTweetNemeton, id, name, description, start, end, &rewards, nil)
}

func makeRPCTask(id, name, description string, start, end time.Time, rewards uint64) Task {
	return makeTask(taskTypeRPC, id, name, description, start, end, &rewards, nil)
}

func makeSnapshotsTask(id, name, description string, start, end time.Time, rewards uint64) Task {
	return makeTask(taskTypeSnapshots, id, name, description, start, end, &rewards, nil)
}

func makeDashboardTask(id, name, description string, start, end time.Time) Task {
	return makeTask(taskTypeDashboard, id, name, description, start, end, nil, nil)
}
