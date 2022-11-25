package tweet

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"okp4/nemeton-leaderboard/app/message"
	"okp4/nemeton-leaderboard/app/offset"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/scheduler"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
)

const ownerOffset = "tweet-search"

type SearchActor struct {
	Hashtag      string
	TweeterToken string
	Client       *http.Client
	Store        *offset.Store
	Context      context.Context
}

func NewSearchActor(eventStore *actor.PID, mongoURI, dbName, tweeterToken string) (*SearchActor, error) {
	ctx := context.Background()
	store, err := offset.NewStore(ctx, mongoURI, dbName, ownerOffset)
	if err != nil {
		return nil, err
	}

	return &SearchActor{
		Hashtag:      "#NemetonOKP4 -is:retweet",
		TweeterToken: tweeterToken,
		Client:       http.DefaultClient,
		Store:        store,
		Context:      ctx,
	}, nil
}

func (a *SearchActor) Receive(ctx actor.Context) {
	switch ctx.Message().(type) {
	case *actor.Started:
		log.Info().Msg("💬 Start tweeter search")
		scheduler.NewTimerScheduler(ctx).SendRepeatedly(0, 10*time.Second, ctx.Self(), &message.SearchTweet{})
	case *message.SearchTweet:
		log.Info().Msg("🧙‍ Start looking for tweets")
		a.searchTweets(a.getSinceID(), "", "")
	case *actor.Stopping:
		log.Info().Msg("🛑 Stop tweeter search")
	}
}

// searchTweets.
func (a *SearchActor) searchTweets(sinceID, nextToken, initialNewestID string) {
	tweets, err := a.fetchTweets(sinceID, nextToken)
	if err != nil {
		log.Error().Err(err).Msg("❌ Failed fetch tweet from twitter API.")
		return
	}

	if tweets.Data == nil {
		log.Info().Msg("🐪 No new tweet found this time, try again later")
		return
	}

	a.handleTweets(tweets)

	if tweets.Meta.NextToken != "" {
		// There is another page, request it.
		newestID := tweets.Meta.NewestID
		if initialNewestID != "" {
			// Keep the initial newest id for save it after pagination
			newestID = initialNewestID
		}
		log.Info().Str("latestTweetId", newestID).Msg("📃 Search tweet on next page")
		a.searchTweets(sinceID, tweets.Meta.NextToken, newestID)
	} else {
		// No new page, save the next sinceId for next scheduled query.
		if initialNewestID != "" {
			if err := a.setSinceID(initialNewestID); err != nil {
				log.Error().Err(err).Msg("❌ Failed save since id value 💾")
				return
			}
		} else {
			if err := a.setSinceID(tweets.Meta.NewestID); err != nil {
				log.Error().Err(err).Msg("❌ Failed save since id value 💾")
				return
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
	query.Add("query", a.Hashtag)
	query.Add("expansions", "author_id")
	query.Add("user.fields", "username")

	if len(nextToken) > 0 {
		query.Add("next_token", nextToken)
	}
	if len(sinceID) > 0 {
		query.Add("since_id", sinceID)
	}

	u.RawQuery = query.Encode()

	request, err := http.NewRequestWithContext(a.Context, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("authorization", fmt.Sprintf("Bearer %s", a.TweeterToken))
	r, err := a.Client.Do(request)
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

func (a *SearchActor) handleTweets(tweets *Response) {
	// TODO: Parse response to register tweet event
}

func (a *SearchActor) getSinceID() string {
	value, _ := a.Store.Get(a.Context)
	switch sinceID := value.(type) {
	case string:
		return sinceID
	default:
		return ""
	}
}

func (a *SearchActor) setSinceID(sinceID string) error {
	return a.Store.Save(a.Context, sinceID)
}
