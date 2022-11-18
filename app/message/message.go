package message

import "okp4/nemeton-leaderboard/app/event"

type PublishEventMessage struct {
	Event event.Event
}
