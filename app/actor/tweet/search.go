package tweet

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"okp4/nemeton-leaderboard/app/event"
	"okp4/nemeton-leaderboard/app/message"
	"okp4/nemeton-leaderboard/app/offset"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/scheduler"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
)

const ownerOffset = "tweet-search"

type SearchActor struct {
	twitterAccount string
	twitterToken   string
	client         *http.Client
	eventStore     *actor.PID
	store          *offset.Store
	context        context.Context
}

func NewSearchActor(eventStore *actor.PID, mongoURI, dbName, twitterToken, twitterAccount string) (*SearchActor, error) {
	ctx := context.Background()
	store, err := offset.NewStore(ctx, mongoURI, dbName, ownerOffset)
	if err != nil {
		return nil, err
	}

	return &SearchActor{
		twitterAccount: twitterAccount,
		twitterToken:   twitterToken,
		client:         http.DefaultClient,
		eventStore:     eventStore,
		store:          store,
		context:        ctx,
	}, nil
}

func (a *SearchActor) Receive(ctx actor.Context) {
	switch ctx.Message().(type) {
	case *actor.Started:
		log.Info().Msg("💬 Start twitter search")
		scheduler.NewTimerScheduler(ctx).SendRepeatedly(0, 10*time.Second, ctx.Self(), &message.SearchTweet{})
	case *message.SearchTweet:
		log.Info().Msg("🧙‍ Start looking for tweets")
		a.searchTweets(ctx, a.getSinceID(), "", "")
	case *actor.Stopping:
		log.Info().Msg("🛑 Stop twitter search")
	}
}

// searchTweets.
func (a *SearchActor) searchTweets(ctx actor.Context, sinceID, nextToken, initialNewestID string) {
	tweets, err := a.fetchTweets(sinceID, nextToken)
	if err != nil {
		log.Error().Err(err).Msg("❌ Failed fetch tweet from twitter API.")
		return
	}

	if tweets.Data == nil {
		log.Info().Msg("🐪 No new tweet found this time, try again later")
		return
	}

	a.handleTweets(ctx, tweets)

	// nolint:gocritic
	// could not be converted into switch statement
	if tweets.Meta.NextToken != "" && initialNewestID != "" {
		// There is another page, request it.
		// Keep the initial newest id for save it after pagination
		log.Info().Str("latestTweetId", initialNewestID).Msg("📃 Search tweet on next page")
		a.searchTweets(ctx, sinceID, tweets.Meta.NextToken, initialNewestID)
	} else if tweets.Meta.NextToken != "" {
		// There is another page, request it.
		log.Info().Str("latestTweetId", tweets.Meta.NewestID).Msg("📃 Search tweet on next page")
		a.searchTweets(ctx, sinceID, tweets.Meta.NextToken, tweets.Meta.NewestID)
	} else {
		// No new page, save the next sinceId for next scheduled query.
		switch initialNewestID {
		case "":
			if err := a.setSinceID(tweets.Meta.NewestID); err != nil {
				log.Panic().Err(err).Msg("❌ Failed save since id value 💾")
			}
		default:
			if err := a.setSinceID(initialNewestID); err != nil {
				log.Panic().Err(err).Msg("❌ Failed save since id value 💾")
			}
		}

		log.Info().Str("latestTweetId", a.getSinceID()).Msg("📃 No new page on tweet search. Looking next time for new tweet.")
	}
}

func (a *SearchActor) fetchTweets(sinceID, nextToken string) (*Response, error) {
	u, err := url.Parse("https://api.twitter.com/2/tweets/search/recent")
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Add("query", fmt.Sprintf("%s -is:retweet", a.twitterAccount))
	query.Add("expansions", "author_id")
	query.Add("user.fields", "username")

	if len(nextToken) > 0 {
		query.Add("next_token", nextToken)
	}
	if len(sinceID) > 0 {
		query.Add("since_id", sinceID)
	}

	u.RawQuery = query.Encode()

	request, err := http.NewRequestWithContext(a.context, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("authorization", fmt.Sprintf("Bearer %s", a.twitterToken))
	r, err := a.client.Do(request)
	if err != nil {
		return nil, err
	}

	defer r.Body.Close()

	if r.StatusCode != 200 {
		return nil, fmt.Errorf("wrong response from twitter api. Status code : %d", r.StatusCode)
	}

	var response Response
	err = json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (a *SearchActor) handleTweets(ctx actor.Context, tweets *Response) {
	for _, tweet := range tweets.Data {
		var author User
		for _, user := range tweets.Includes.Users {
			if user.ID == tweet.AuthorID {
				author = user
				break
			}
		}

		e := NewTweetEvent{
			ID:       tweet.ID,
			AuthorID: tweet.AuthorID,
			Text:     tweet.Text,
			User:     author,
		}

		eventData, err := e.Marshall()
		if err != nil {
			log.Err(err).Msg("❌ Failed to marshall event to map interface")
			return
		}
		ctx.Send(a.eventStore, &message.PublishEventMessage{Event: event.NewEvent(NewTweetEventType, eventData)})
	}
}

func (a *SearchActor) getSinceID() string {
	value, _ := a.store.Get(a.context)
	switch sinceID := value.(type) {
	case string:
		return sinceID
	default:
		return ""
	}
}

func (a *SearchActor) setSinceID(sinceID string) error {
	return a.store.Save(a.context, sinceID)
}
