package event

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Stream struct {
	closed *atomic.Bool
	evtCh  chan Event
	errCh  chan error
	wg     *sync.WaitGroup

	crtErr   error
	crtEvent *Event
}

func (s *Stream) Next() bool {
	select {
	case evt, ok := <-s.evtCh:
		if !ok {
			return false
		}
		s.crtEvent = &evt
		return true
	case err, ok := <-s.errCh:
		if ok {
			s.crtErr = err
		}
		return false
	}
}

func (s *Stream) Err() error {
	return s.crtErr
}

func (s *Stream) Event() *Event {
	return s.crtEvent
}

func (s *Stream) Close() {
	s.closed.Store(true)
	close(s.evtCh)
	close(s.errCh)
	s.wg.Wait()
}

func openStream(ctx context.Context, col *mongo.Collection, from *primitive.ObjectID) (*Stream, error) {
	watch, catchup, err := fetch(ctx, col, from)
	if err != nil {
		return nil, err
	}

	stream := &Stream{
		closed: &atomic.Bool{},
		evtCh:  make(chan Event, 100),
		errCh:  make(chan error, 1),
		wg:     &sync.WaitGroup{},
	}

	go stream.start(ctx, watch, catchup)
	return stream, nil
}

func (s *Stream) start(ctx context.Context, watch *mongo.ChangeStream, catchUp *mongo.Cursor) {
	s.wg.Add(1)
	defer func() {
		_ = watch.Close(ctx)
		s.wg.Done()
	}()

	caughtUpIDs, err := s.catchUp(ctx, catchUp)
	if err != nil {
		s.errCh <- err
		return
	}

	for {
		if s.closed.Load() {
			return
		}

		evt, err := s.readWatch(ctx, watch, caughtUpIDs)
		if err != nil {
			s.errCh <- err
			return
		}
		if evt != nil {
			s.evtCh <- *evt
		} else {
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func (s *Stream) readWatch(
	ctx context.Context,
	watch *mongo.ChangeStream,
	idsToIgnore map[primitive.ObjectID]interface{},
) (*Event, error) {
	if !watch.TryNext(ctx) {
		return nil, watch.Err()
	}

	var res struct {
		OperationType string `bson:"operationType"`
		FullDocument  Event  `bson:"fullDocument"`
	}
	if err := watch.Decode(&res); err != nil {
		return nil, err
	}

	if res.OperationType != "insert" {
		return s.readWatch(ctx, watch, idsToIgnore)
	}

	evt := res.FullDocument
	if _, ok := idsToIgnore[evt.ID]; ok {
		delete(idsToIgnore, evt.ID)
		return s.readWatch(ctx, watch, idsToIgnore)
	}

	return &evt, nil
}

func (s *Stream) catchUp(ctx context.Context, c *mongo.Cursor) (map[primitive.ObjectID]interface{}, error) {
	defer func() {
		_ = c.Close(ctx)
	}()

	ids := make(map[primitive.ObjectID]interface{})
	for c.Next(ctx) {
		var evt Event
		if err := c.Decode(&evt); err != nil {
			return ids, err
		}

		ids[evt.ID] = nil
		s.evtCh <- evt

		if s.closed.Load() {
			return ids, nil
		}
	}
	return ids, nil
}

func fetch(ctx context.Context, col *mongo.Collection, from *primitive.ObjectID) (*mongo.ChangeStream, *mongo.Cursor, error) {
	watch, err := col.Watch(ctx, mongo.Pipeline{}, options.ChangeStream())
	if err != nil {
		return nil, nil, err
	}

	var filter bson.M
	if from != nil {
		filter = bson.M{
			"_id": bson.M{
				"$gt": from,
			},
		}
	}
	catchUp, err := col.
		Find(
			ctx,
			filter,
			&options.FindOptions{
				Sort: bson.M{
					"_id": 1,
				},
			},
		)
	if err != nil {
		_ = watch.Close(ctx)
		return nil, nil, err
	}

	return watch, catchUp, nil
}
