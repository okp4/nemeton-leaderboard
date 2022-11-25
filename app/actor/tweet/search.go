package tweet

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"okp4/nemeton-leaderboard/app/message"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/scheduler"
	"github.com/rs/zerolog/log"
)

type SearchActor struct {
	Hashtag      string
	TweeterToken string
	Client       *http.Client
	SinceId      string
}

func NewSearchActor(eventStore *actor.PID, mongoURI, dbName, tweeterToken string) (*SearchActor, error) {
	return &SearchActor{
		Hashtag:      "#NemetonOKP4 -is:retweet",
		TweeterToken: tweeterToken,
		Client:       http.DefaultClient,
		SinceId:      "",
	}, nil
}

func (a *SearchActor) Receive(ctx actor.Context) {
	switch ctx.Message().(type) {
	case *actor.Started:
		log.Info().Msg("💬 Start tweeter search")
		scheduler.NewTimerScheduler(ctx).SendRepeatedly(0, 10*time.Second, ctx.Self(), &message.SearchTweet{})
	case *message.SearchTweet:
		log.Info().Msg("🧙‍ Start looking for tweets")
		a.searchTweets(a.SinceId, "", "")
	case *actor.Stopping:
		log.Info().Msg("🛑 Stop tweeter search")
	}
}

// searchTweets
func (a *SearchActor) searchTweets(sinceId, nextToken, initialNewestId string) {
	tweets, err := a.fetchTweets(sinceId, nextToken)
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
		newestId := tweets.Meta.NewestId
		if initialNewestId != "" {
			// Keep the initial newest id for save it after pagination
			newestId = initialNewestId
		}
		log.Info().Str("latestTweetId", newestId).Msg("📃 Search tweet on next page")
		a.searchTweets(sinceId, tweets.Meta.NextToken, newestId)
	} else {
		// No new page, save the next sinceId for next scheduled query.
		if initialNewestId != "" {
			a.SinceId = initialNewestId
		} else {
			a.SinceId = tweets.Meta.NewestId
		}
		log.Info().Str("latestTweetId", a.SinceId).Msg("📃 No new page on tweet search. Looking next time for new tweet.")
	}
}

func (a *SearchActor) fetchTweets(sinceId, nextToken string) (*Response, error) {
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
	if len(sinceId) > 0 {
		query.Add("since_id", sinceId)
	}

	u.RawQuery = query.Encode()

	request, err := http.NewRequest(http.MethodGet, u.String(), nil)
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
