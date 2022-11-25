package tweet

type Response struct {
	Data []struct {
		Id                  string   `json:"id"`
		EditHistoryTweetIds []string `json:"edit_history_tweet_ids"`
		AuthorId            string   `json:"author_id"`
		Text                string   `json:"text"`
	} `json:"data"`
	Includes struct {
		Users []struct {
			Username    string `json:"username"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Id          string `json:"id"`
		} `json:"users"`
	} `json:"includes"`
	Meta struct {
		NewestId    string `json:"newest_id"`
		OldestId    string `json:"oldest_id"`
		ResultCount int    `json:"result_count"`
		NextToken   string `json:"next_token"`
	} `json:"meta"`
}
