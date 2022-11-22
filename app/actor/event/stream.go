package event

import (
	"sync"

	"okp4/nemeton-leaderboard/app/message"

	"okp4/nemeton-leaderboard/app/event"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/rs/zerolog/log"
)

type StreamHandlerActor struct {
	dst    *actor.PID
	wg     *sync.WaitGroup
	stream *event.Stream
}

func NewStreamHandlerActor(stream *event.Stream, dst *actor.PID) *StreamHandlerActor {
	return &StreamHandlerActor{
		dst:    dst,
		wg:     &sync.WaitGroup{},
		stream: stream,
	}
}

func (a *StreamHandlerActor) Receive(ctx actor.Context) {
	switch ctx.Message().(type) {
	case *actor.Started:
		a.handleStart(ctx)
	case *actor.Stopping:
		a.handleStop()
	}
}

func (a *StreamHandlerActor) handleStart(ctx actor.Context) {
	a.wg = &sync.WaitGroup{}
	a.wg.Add(1)

	go a.processStream(ctx)
}

func (a *StreamHandlerActor) handleStop() {
	a.stream.Close()
	a.wg.Wait()
}

func (a *StreamHandlerActor) processStream(ctx actor.Context) {
	defer a.wg.Done()

	for a.stream.Next() {
		ctx.Send(a.dst, &message.NewEventMessage{Event: *a.stream.Event()})
	}

	if err := a.stream.Err(); err != nil {
		log.Err(a.stream.Err()).Msg("❌ stream stopped unexpectedly")
	}
	ctx.Send(a.dst, &message.BrokenStreamMessage{})
	ctx.Stop(ctx.Self())
}
