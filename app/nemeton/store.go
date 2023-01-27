package nemeton

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"time"

	"okp4/nemeton-leaderboard/app/util"

	"github.com/cosmos/cosmos-sdk/types"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	phasesCollectionName     = "phases"
	validatorsCollectionName = "validators"
)

var omitMissedBlocks = bson.M{"missedBlocks": 0}

type Store struct {
	db     *mongo.Database
	phases []*Phase
}

func NewStore(ctx context.Context, mongoURI, dbName string) (*Store, error) {
	db, err := util.OpenMongoDatabase(ctx, mongoURI, dbName)
	if err != nil {
		return nil, err
	}

	store := &Store{
		db: db,
	}

	if err := store.init(ctx); err != nil {
		return nil, err
	}

	return store, nil
}

func (s *Store) Close(ctx context.Context) error {
	return s.db.Client().Disconnect(ctx)
}

func (s *Store) init(ctx context.Context) error {
	phases := s.db.Collection(phasesCollectionName)
	count, err := phases.CountDocuments(ctx, bson.M{})
	if err != nil {
		return err
	}

	if count == 0 {
		for _, phase := range bootstrapPhases() {
			_, err := phases.InsertOne(ctx, phase)
			if err != nil {
				return err
			}
		}
	}

	c, err := phases.Find(ctx, bson.M{})
	if err != nil {
		return err
	}

	for c.Next(ctx) {
		var phase Phase
		if err := c.Decode(&phase); err != nil {
			return err
		}
		s.phases = append(s.phases, &phase)
	}

	if err := s.ensureIndexes(ctx); err != nil {
		return err
	}
	return nil
}

func (s *Store) ensureIndexes(ctx context.Context) error {
	_, err := s.db.Collection(validatorsCollectionName).
		Indexes().
		CreateMany(
			ctx,
			[]mongo.IndexModel{
				{Keys: bson.M{"points": 1}},
				{Keys: bson.M{"moniker": 1}},
				{Keys: bson.M{"valoper": 1}, Options: options.Index().SetUnique(true)},
				{Keys: bson.M{"delegator": 1}, Options: options.Index().SetUnique(true)},
				{Keys: bson.M{"valcons": 1}, Options: options.Index().SetUnique(true)},
				{Keys: bson.M{"twitter": 1}, Options: options.Index().SetSparse(true).SetUnique(true)},
				{Keys: bson.M{"discord": 1}, Options: options.Index().SetUnique(true)},
			},
		)
	return err
}

func (s *Store) GetPhase(number int) *Phase {
	for _, phase := range s.phases {
		if phase.Number == number {
			return phase
		}
	}
	return nil
}

// GetCurrentPhase returns the phase which is in progress, if any.
// WARNING: Do not use this when processing and event, the event sourcing shall use only the event's context: we need
// the phase in progress at the time of the event, not now.
func (s *Store) GetCurrentPhase() *Phase {
	return s.GetCurrentPhaseAt(time.Now())
}

func (s *Store) GetCurrentPhaseAt(at time.Time) *Phase {
	for _, phase := range s.phases {
		if phase.InProgressAt(at) {
			return phase
		}
	}
	return nil
}

func (s *Store) GetAllPhases() []*Phase {
	return s.phases
}

func (s *Store) GetUnstartedPhases() []*Phase {
	return s.GetPhases(func(p Phase) bool {
		return !p.Started()
	})
}

func (s *Store) GetFinishedPhases() []*Phase {
	return s.GetPhases(func(p Phase) bool {
		return p.Finished()
	})
}

func (s *Store) GetPhases(criteriaFn func(p Phase) bool) []*Phase {
	var filtered []*Phase
	for _, phase := range s.phases {
		if criteriaFn(*phase) {
			filtered = append(filtered, phase)
		}
	}
	return filtered
}

func (s *Store) GetValidatorByCursor(ctx context.Context, c Cursor) (*Validator, error) {
	return s.GetValidatorBy(ctx, bson.M{"_id": c.objectID})
}

func (s *Store) GetValidatorByValoper(ctx context.Context, addr types.ValAddress) (*Validator, error) {
	return s.GetValidatorBy(ctx, bson.M{"valoper": addr})
}

func (s *Store) GetValidatorByDelegator(ctx context.Context, addr types.AccAddress) (*Validator, error) {
	return s.GetValidatorBy(ctx, bson.M{"delegator": addr})
}

func (s *Store) GetValidatorByDiscord(ctx context.Context, discord string) (*Validator, error) {
	return s.GetValidatorBy(ctx, bson.M{"discord": discord})
}

func (s *Store) GetValidatorByTwitter(ctx context.Context, twitter string) (*Validator, error) {
	return s.GetValidatorBy(ctx, bson.M{"twitter": twitter})
}

func (s *Store) GetValidatorBy(ctx context.Context, filter bson.M) (*Validator, error) {
	res := s.db.Collection(validatorsCollectionName).
		FindOne(ctx, filter, options.FindOne().SetProjection(omitMissedBlocks))
	if err := res.Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	var val Validator
	return &val, res.Decode(&val)
}

func (s *Store) GetValidatorRank(ctx context.Context, cursor Cursor) (int, error) {
	count, err := s.db.Collection(validatorsCollectionName).
		CountDocuments(
			ctx,
			bson.M{
				"$or": bson.A{
					bson.M{
						"points": bson.M{"$gt": cursor.points},
					},
					bson.M{
						"points": cursor.points,
						"_id":    bson.M{"$lt": cursor.objectID},
					},
				},
			},
		)

	return int(count) + 1, err
}

func (s *Store) CountValidators(ctx context.Context) (int64, error) {
	return s.db.Collection(validatorsCollectionName).CountDocuments(ctx, bson.M{})
}

func (s *Store) GetBoard(ctx context.Context, search *string, limit int, after *Cursor) ([]*Validator, bool, error) {
	c, err := s.db.Collection(validatorsCollectionName).Find(
		ctx,
		makeBoardFilter(search, after),
		options.Find().
			SetSort(
				bson.D{
					bson.E{Key: "points", Value: -1},
					bson.E{Key: "_id", Value: 1},
				},
			).
			SetLimit(int64(limit+1)),
		options.Find().SetProjection(omitMissedBlocks),
	)
	if err != nil {
		return nil, false, err
	}
	defer func() {
		_ = c.Close(ctx)
	}()

	validators := make([]*Validator, 0, limit)
	count := 0
	for count < limit && c.Next(ctx) {
		count++
		var validator Validator
		if err := c.Decode(&validator); err != nil {
			return nil, false, err
		}
		validators = append(validators, &validator)
	}

	return validators, c.Next(ctx), nil
}

func makeBoardFilter(search *string, after *Cursor) bson.M {
	filters := bson.A{}
	if after != nil {
		filters = append(filters, bson.M{
			"$or": bson.A{
				bson.M{
					"points": bson.M{"$lt": after.points},
				},
				bson.M{
					"points": after.points,
					"_id":    bson.M{"$gt": after.objectID},
				},
			},
		})
	}
	if search != nil {
		filters = append(filters, bson.M{
			"$or": bson.A{
				bson.M{"moniker": bson.M{"$regex": fmt.Sprintf(".*%s.*", *search), "$options": "i"}},
				bson.M{"valoper": bson.M{"$regex": fmt.Sprintf(".*%s.*", *search), "$options": "i"}},
			},
		})
	}

	var filter bson.M
	if len(filters) > 0 {
		filter = bson.M{
			"$and": filters,
		}
	}
	return filter
}

func (s *Store) CreateGentxValidator(
	ctx context.Context,
	createdAt time.Time,
	msgCreateVal *stakingtypes.MsgCreateValidator,
	discord, country string,
	twitter *string,
) error {
	validator, err := MakeValidatorFromMsg(msgCreateVal, discord, country, twitter)
	if err != nil {
		return err
	}

	points := uint64(0)
	var tasks map[int]map[string]TaskState
	p, task := s.getTaskPhaseByType(taskTypeGentx, createdAt)
	if p != nil && task != nil {
		points = *task.Rewards
		tasks = map[int]map[string]TaskState{
			p.Number: {
				task.ID: {
					Completed:    true,
					EarnedPoints: *task.Rewards,
				},
			},
		}
	}

	validator.Points = &points
	validator.Tasks = tasks

	_, err = s.db.Collection(validatorsCollectionName).InsertOne(ctx, validator)
	return err
}

func (s *Store) RegisterValidator(
	ctx context.Context,
	valoper types.ValAddress,
	delegator types.AccAddress,
	valcons types.ConsAddress,
	description stakingtypes.Description,
	discord, country string,
	twitter *string,
	lastHeight int64,
) error {
	validator, err := NewValidator(valoper, delegator, valcons, description, discord, country, twitter)
	if err != nil {
		return err
	}

	res, err := s.db.Collection(validatorsCollectionName).InsertOne(ctx, validator)
	if err != nil {
		return err
	}

	_, err = s.db.Collection(validatorsCollectionName).UpdateOne(
		ctx,
		bson.M{"_id": res.InsertedID},
		bson.M{
			"missedBlocks": bson.A{
				bson.M{
					"from": 1,
					"to":   lastHeight + 1,
				},
			},
		},
	)
	return err
}

func (s *Store) UpdateValidator(
	ctx context.Context,
	delegator types.AccAddress,
	valoper types.ValAddress,
	valcons types.ConsAddress,
	description stakingtypes.Description,
	discord, country string,
	twitter *string,
) error {
	validator, err := NewValidator(valoper, delegator, valcons, description, discord, country, twitter)
	if err != nil {
		return err
	}

	_, err = s.db.Collection(validatorsCollectionName).
		UpdateOne(
			ctx,
			bson.M{
				"delegator": delegator,
			},
			bson.M{
				"$set": validator,
			},
		)
	return err
}

func (s *Store) RemoveValidator(ctx context.Context, valoper types.ValAddress) error {
	_, err := s.db.Collection(validatorsCollectionName).
		DeleteOne(
			ctx,
			bson.M{
				"valoper": valoper,
			})
	return err
}

func (s *Store) RegisterValidatorURL(ctx context.Context,
	when time.Time,
	urlType string,
	validator types.ValAddress,
	url *url.URL,
	rewards *uint64,
) error {
	var field string
	switch urlType {
	case TaskTypeRPC:
		field = "rpcEndpoint"
	case TaskTypeDashboard:
		field = "dashboard"
	case TaskTypeSnapshots:
		field = "snapshot"
	}

	filter := bson.M{"valoper": validator}
	_, err := s.db.Collection(validatorsCollectionName).UpdateOne(ctx,
		filter,
		bson.M{"$set": bson.M{field: url}},
	)
	if err != nil {
		return err
	}

	if phase, task := s.getTaskPhaseByType(urlType, when); phase != nil && task != nil {
		var r uint64
		switch urlType {
		case TaskTypeRPC, TaskTypeSnapshots:
			r = *task.Rewards
		case TaskTypeDashboard:
			r = *rewards
		}
		return s.ensureTaskCompleted(ctx, filter, phase.Number, task.ID, r)
	}
	return fmt.Errorf("could not find corresponding phase and task at %s. Did this task begun ? ", when.Format(time.RFC3339))
}

func (s *Store) UpdateValidatorUptime(ctx context.Context, consensusAddrs []types.ConsAddress, height int64) error {
	model := []mongo.WriteModel{
		mongo.NewUpdateManyModel().
			SetFilter(
				bson.M{
					"$and": bson.A{
						bson.M{"valcons": bson.M{"$nin": consensusAddrs}},
						bson.M{"missedBlocks.to": bson.M{"$eq": height}},
					},
				}).
			SetUpdate(
				bson.M{"$inc": bson.M{"missedBlocks.$[down].to": 1}},
			).
			SetArrayFilters(options.ArrayFilters{
				Filters: bson.A{bson.M{"down.to": bson.M{"$eq": height}}},
			}),
		mongo.NewUpdateManyModel().
			SetFilter(
				bson.M{
					"$and": bson.A{
						bson.M{"valcons": bson.M{"$nin": consensusAddrs}},
						bson.M{"$or": bson.A{
							bson.M{"missedBlocks": bson.M{"$size": 0}},
							bson.M{"missedBlocks.to": bson.M{"$not": bson.M{"$gte": height}}},
						}},
					},
				}).
			SetUpdate(
				bson.M{
					"$push": bson.M{"missedBlocks": bson.M{"from": height, "to": height + 1}},
				},
			),
	}
	opts := options.BulkWrite().SetOrdered(true)
	_, err := s.db.Collection(validatorsCollectionName).BulkWrite(ctx, model, opts)
	return err
}

func (s *Store) CompleteTweetTask(ctx context.Context, when time.Time, username string, phase *Phase, task Task) error {
	filter := bson.M{"twitter": username}
	if !phase.InProgressAt(when) || !task.InProgressAt(when) {
		return fmt.Errorf("could not complete task since task or phase is not in progress at %s", when.Format(time.RFC3339))
	}

	return s.ensureTaskCompleted(ctx, filter, phase.Number, task.ID, *task.Rewards)
}

func (s *Store) CompleteNodeSetupTask(ctx context.Context, when time.Time, vals []types.ConsAddress) error {
	phase, task := s.getTaskPhaseByType(taskTypeNodeSetup, when)
	if phase != nil && task != nil {
		return s.ensureTaskCompleted(ctx, bson.M{"valcons": bson.M{"$in": vals}}, phase.Number, task.ID, *task.Rewards)
	}
	return nil
}

func (s *Store) ManualSubmitTask(
	ctx context.Context,
	valoper types.ValAddress,
	phaseNB int,
	taskID string,
) error {
	phase := s.GetPhase(phaseNB)
	var task *Task
	for i, it := range phase.Tasks {
		if it.ID == taskID {
			task = &phase.Tasks[i]
		}
	}
	if task == nil {
		return fmt.Errorf("task '%s' not found in phase '%d'", taskID, phaseNB)
	}

	if !task.WithSubmission() {
		return fmt.Errorf("task '%s' in phase '%d' has no submission", taskID, phaseNB)
	}

	_, err := s.db.Collection(validatorsCollectionName).UpdateOne(ctx,
		bson.M{
			"valoper": valoper,
		},
		bson.M{
			"$set": bson.M{
				fmt.Sprintf("tasks.%d.%s.submitted", phaseNB, taskID): true,
			},
		})
	return err
}

func (s *Store) ManualCompleteTask(
	ctx context.Context,
	valoper types.ValAddress,
	phaseNB int,
	taskID string,
	rewards *uint64,
) error {
	phase := s.GetPhase(phaseNB)
	var task *Task
	for i, it := range phase.Tasks {
		if it.ID == taskID {
			task = &phase.Tasks[i]
		}
	}
	if task == nil {
		return fmt.Errorf("task '%s' not found in phase '%d'", taskID, phaseNB)
	}

	points := uint64(0)
	if rewards != nil {
		points = *rewards
	} else {
		if task.Rewards == nil {
			return fmt.Errorf("no rewards found for task '%s' in phase '%d'", taskID, phaseNB)
		}
		points = *task.Rewards
	}

	return s.ensureTaskCompleted(ctx, bson.M{"valoper": valoper}, phaseNB, taskID, points)
}

func (s *Store) ensureTaskCompleted(ctx context.Context, filter bson.M, phase int, task string, rewards uint64) error {
	_, err := s.db.Collection(validatorsCollectionName).UpdateMany(ctx,
		bson.M{
			"$and": bson.A{
				filter,
				bson.M{fmt.Sprintf("tasks.%d.%s.completed", phase, task): bson.M{"$ne": true}},
			},
		},
		bson.M{
			"$set": bson.M{
				fmt.Sprintf("tasks.%d.%s.completed", phase, task): true,
				fmt.Sprintf("tasks.%d.%s.points", phase, task):    rewards,
			},
			"$inc": bson.M{
				"points": rewards,
			},
		})
	return err
}

// getTaskPhaseByType return the current phase at the given time and the current **first** task for the given task type.
func (s *Store) getTaskPhaseByType(id string, at time.Time) (*Phase, *Task) {
	if phase, tasks := s.getTasksPhaseByType(id, at); len(tasks) > 0 {
		return phase, tasks[0]
	} else {
		return phase, nil
	}
}

// getTasksPhaseByType return the current phase at the given time and all tasks for the given task type.
func (s *Store) getTasksPhaseByType(id string, at time.Time) (phase *Phase, tasks []*Task) {
	phase = s.GetCurrentPhaseAt(at)
	tasks = make([]*Task, 0)
	if phase != nil {
		for _, task := range phase.Tasks {
			if task.Type == id && task.InProgressAt(at) {
				tasks = append(tasks, &task)
			}
		}
	}
	return phase, tasks
}

func (s *Store) UpdatePhaseBlocks(ctx context.Context, blockTime time.Time, height int64) error {
	_, err := s.db.Collection(phasesCollectionName).UpdateOne(ctx, bson.M{
		"startDate": bson.M{"$lte": blockTime},
		"endDate":   bson.M{"$gt": blockTime},
	}, bson.A{
		bson.M{
			"$set": bson.M{
				"blocks": bson.M{
					"$ifNull": bson.A{
						"$blocks",
						bson.M{"from": height, "to": height + 1},
						"$blocks",
					},
				},
			},
		},
		bson.M{"$set": bson.M{"blocks.to": height + 1}},
	})
	return err
}

func (s *Store) GetPhaseBlocks(ctx context.Context, number int) (*BlockRange, error) {
	var phase struct {
		Blocks *BlockRange `bson:"blocks"`
	}
	if err := s.db.Collection(phasesCollectionName).FindOne(ctx,
		bson.M{"_id": number},
		options.FindOne().SetProjection(bson.M{"blocks": 1}),
	).Decode(&phase); err != nil {
		return nil, err
	}

	return phase.Blocks, nil
}

func (s *Store) GetPreviousPhaseByBlock(ctx context.Context, height int64) (*Phase, error) {
	res := s.db.Collection(phasesCollectionName).
		FindOne(ctx, bson.M{"blocks.to": height}, options.FindOne())
	if err := res.Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	var phase Phase
	return &phase, res.Decode(&phase)
}

// CompleteValidatorsUptimeForPhase is used to concat all missed blocks on a given phase and calculate the number of
// points rewarded.
//
//nolint:funlen
func (s *Store) CompleteValidatorsUptimeForPhase(ctx context.Context, phase *Phase) error {
	// Retrieve task at the phase end time since event time is thrown at the next phase (so task and phase are ended)
	_, task := s.getTaskPhaseByType(taskTypeUptime, phase.EndDate)
	if task == nil {
		return fmt.Errorf("could not retrieve uptime task for phase %d", phase.Number)
	}

	reward := task.GetParamMaxPoints()
	if reward == nil {
		return fmt.Errorf("could not retrieve the maximum number of point for task %s", task.ID)
	}

	_, err := s.db.Collection(validatorsCollectionName).Aggregate(ctx, bson.A{
		bson.M{"$lookup": bson.M{
			"from": "phases",
			"pipeline": bson.A{
				bson.M{"$match": bson.M{"_id": phase.Number}},
				bson.M{"$project": bson.M{"blocks": 1}},
			},
			"as": "currentPhase",
		}},
		bson.M{"$unwind": bson.M{"path": "$missedBlocks", "preserveNullAndEmptyArrays": true}},
		bson.M{"$unwind": bson.M{"path": "$currentPhase"}},
		bson.M{"$replaceRoot": bson.M{
			"newRoot": bson.M{
				"count": bson.M{
					"$switch": bson.M{
						"branches": bson.A{
							bson.M{
								"case": bson.M{"$and": bson.A{
									bson.M{"$gte": bson.A{"$missedBlocks.from", "$currentPhase.blocks.from"}},
									bson.M{"$lte": bson.A{"$missedBlocks.to", "$currentPhase.blocks.to"}},
								}},
								"then": bson.M{"$subtract": bson.A{"$missedBlocks.to", "$missedBlocks.from"}},
							},
							bson.M{
								"case": bson.M{"$and": bson.A{
									bson.M{"$lt": bson.A{"$missedBlocks.from", "$currentPhase.blocks.from"}},
									bson.M{"$lte": bson.A{"$missedBlocks.to", "$currentPhase.blocks.to"}},
									bson.M{"$gte": bson.A{"$missedBlocks.to", "$currentPhase.blocks.from"}},
								}},
								"then": bson.M{"$subtract": bson.A{"$missedBlocks.to", "$currentPhase.blocks.from"}},
							},
							bson.M{
								"case": bson.M{"$and": bson.A{
									bson.M{"$gte": bson.A{"$missedBlocks.from", "$currentPhase.blocks.from"}},
									bson.M{"$gt": bson.A{"$missedBlocks.to", "$currentPhase.blocks.to"}},
									bson.M{"$lte": bson.A{"$missedBlocks.from", "$currentPhase.blocks.to"}},
								}},
								"then": bson.M{"$subtract": bson.A{"$currentPhase.blocks.to", "$missedBlocks.from"}},
							},
						},
						"default": 0,
					},
				},
				"currentPhase": "$currentPhase",
				"validator":    "$_id",
				"points":       "$points",
			},
		}},
		bson.M{"$group": bson.M{
			"_id":              "$validator",
			"totalMissedBlock": bson.M{"$sum": "$count"},
			"phase":            bson.M{"$mergeObjects": "$currentPhase"},
			"points":           bson.M{"$first": "$points"},
		}},
		bson.M{"$addFields": bson.M{
			"uptime": bson.M{"$subtract": bson.A{
				bson.M{
					"$pow": bson.A{*reward + 1, bson.M{
						"$multiply": bson.A{0.01, bson.M{
							"$subtract": bson.A{100, bson.M{"$divide": bson.A{
								bson.M{"$multiply": bson.A{100, "$totalMissedBlock"}},
								bson.M{"$subtract": bson.A{"$phase.blocks.to", "$phase.blocks.from"}},
							}}},
						}},
					}},
				},
				1,
			}},
		}},
		bson.M{"$addFields": bson.M{
			"newPoints": bson.M{"$add": bson.A{"$points", "$uptime"}},
		}},
		bson.M{"$merge": bson.M{
			"into": "validators",
			"on":   "_id",
			"let":  bson.M{"uptime": "$uptime", "newPoints": "$newPoints"},
			"whenMatched": bson.A{
				bson.M{
					"$set": bson.M{
						fmt.Sprintf("tasks.%d.%s.points", phase.Number, task.ID):    bson.M{"$toLong": "$$uptime"},
						fmt.Sprintf("tasks.%d.%s.completed", phase.Number, task.ID): true,
						"points": bson.M{"$toLong": "$$newPoints"},
					},
				},
			},
			"whenNotMatched": "discard",
		}},
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) CompleteVoteProposalTask(ctx context.Context, when time.Time, msgVotes []v1.MsgVote) error {
	if len(msgVotes) == 0 {
		return nil
	}

	phase, tasks := s.getTasksPhaseByType(TaskTypeVoteProposal, when)
	if phase == nil || tasks == nil || len(tasks) == 0 {
		return nil
	}

	for _, task := range tasks {
		proposalID := task.GetParamProposalID()
		if proposalID == nil {
			return fmt.Errorf("could not retrieve linked proposal ID for task %s", task.ID)
		}

		addrs := make([]types.AccAddress, 0, len(msgVotes))
		for _, vote := range msgVotes {
			if vote.ProposalId == *proposalID {
				addr, err := types.AccAddressFromBech32(vote.Voter)
				if err != nil {
					return err
				}
				addrs = append(addrs, addr)
			}
		}
		return s.ensureTaskCompleted(ctx, bson.M{"delegator": bson.M{"$in": addrs}}, phase.Number, task.ID, *task.Rewards)
	}

	return nil
}
