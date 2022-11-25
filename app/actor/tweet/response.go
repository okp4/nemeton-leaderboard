package tweet

type Response struct {
	Data []struct {
		ID                  string   `json:"id"`
		EditHistoryTweetIds []string `json:"edit_history_tweet_ids"`
		AuthorID            string   `json:"author_id"`
		Text                string   `json:"text"`
	} `json:"data"`
	Includes struct {
		Users []struct {
			Username    string `json:"username"`
			Name        string `json:"name"`
			Description string `json:"description"`
			ID          string `json:"id"`
		} `json:"users"`
	} `json:"includes"`
	Meta struct {
		NewestID    string `json:"newest_id"`
		OldestID    string `json:"oldest_id"`
		ResultCount int    `json:"result_count"`
		NextToken   string `json:"next_token"`
	} `json:"meta"`
}
