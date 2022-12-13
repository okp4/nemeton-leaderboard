package nemeton

import "time"

// nolint: funlen,lll
func bootstrapPhases() []Phase {
	r1000 := uint64(1000)
	r2000 := uint64(2000)
	r500 := uint64(500)
	return []Phase{
		{
			Number:      1,
			Name:        "Sidh",
			Description: "This first phase is pretty basic, it is dedicated to setting up Druids' validator environment, participating in the genesis, and getting familiar with the OKP4 testnet.",
			StartDate:   time.Date(2022, time.December, 1, 11, 0, 0, 0, time.UTC),
			EndDate:     time.Date(2023, time.January, 1, 23, 59, 59, 0, time.UTC),
			Tasks: []Task{
				{
					ID:   "0",
					Type: taskTypeGentx,
					Name: "Submit your gentx",
					Description: `When?
From Dec 1st to Dec 12th.

Description
Before starting the network we need to register your validator in the genesis.json.

The gentx creation and registration procedure are detailed here: https://github.com/okp4/networks/tree/main/chains/nemeton-1.

Your gentx shall be submitted through an issue on the https://github.com/okp4/networks/ github repository.

This task is required to make you visible on the Leaderboard.

Rewards
1000 points.

Judging Criteria
You will receive the points once the OKP4 team has integrated your gentx in the genesis.

How to submit
Send the issue number in a private message to Anik#9282 on Discord.`,
					StartDate: time.Date(2022, time.December, 1, 11, 0, 0, 0, time.UTC),
					EndDate:   time.Date(2022, time.December, 13, 0, 0, 0, 0, time.UTC),
					Rewards:   &r1000,
				},
				{
					ID:   "1",
					Type: taskTypeBasic,
					Name: "Setup your node",
					Description: `When?
From Dec 14th at 15pm UTC to Jan 1st.

Description
It is time to make the okp4-nemeton-1 network alive, you have to setup your node and join the network. The technical documentation regarding node setup and network join information is here: https://docs.okp4.network/nodes/introduction.

Rewards
2000 points.

Judging Criteria
Your validator is up and running.

How to submit
The validator presence in the consensus will be automatically checked.`,
					StartDate: time.Date(2022, time.December, 14, 15, 0, 0, 0, time.UTC),
					EndDate:   time.Date(2023, time.January, 1, 23, 59, 59, 0, time.UTC),
					Rewards:   &r2000,
				},
				{
					ID:   "2",
					Type: taskTypeTweetNemeton,
					Name: "Tweet about the OKP4 testnet",
					Description: `When?
From Dec 12th to Jan 1st.
No rush to tweet about it when the task opens – it is better to spread them across that period.

Description
Publish a tweet about the Nemeton testnet while including the @okp4_protocol tag using your validator twitter account.
Feel free to share your excitement!

Rewards
500 points.

Judging Criteria
You will receive the points once the OKP4 team has reviewed your tweet.

How to submit
Tweets are automatically tracked.`,
					StartDate: time.Date(2022, time.December, 12, 0, 0, 0, 0, time.UTC),
					EndDate:   time.Date(2023, time.January, 1, 23, 59, 59, 0, time.UTC),
					Rewards:   &r500,
				},
				{
					ID:   "3",
					Type: taskTypeUptime,
					Name: "Uptime",
					Description: `When?
From Dec 14th at 15pm UTC to Jan 1st.

Description
Maintain the best uptime with your validator.

Rewards
Up to 2500 points with the following formula: 2501^0,01x - 1 with x = %uptime.

Judging Criteria
The less blocks your validator miss, the more points you get.

How to submit
Missed blocks are automatically tracked.`,
					StartDate: time.Date(2022, time.December, 14, 16, 0, 0, 0, time.UTC),
					EndDate:   time.Date(2023, time.January, 1, 23, 59, 59, 0, time.UTC),
				},
				{
					ID:   "4",
					Type: taskTypeSubmission,
					Name: "Submit an original content related to validation",
					Description: `When?
From Dec 12th to Jan 1st

Description
Based on your experience as a validator, write an original article, twitter thread or video content providing value to other validators and the community in general. Content must be in English.
The content may be used later to bring improvements on the node operator’s docs (https://docs.okp4.network/nodes/introduction), be referenced in OKP4's Medium (https://blog.okp4.network/), or shared in social networks.
If you’ve seen great documentation, articles or content ideas elsewhere, help us bring something similar to OKP4. Feel free to be creative if you’re in the right mood!

Rewards
Up to 10 000 points per druid will be attributed, capped at 150 000 points in total.

Judging criteria
OKP4 team will judge if any submission deserves points or not, based on:

Overall relevance
Originality
Completeness
Readability
Useful tips
Good surprises…
Non-relevant submissions or low-value ones will earn 0 points.

How to submit
Share the content links to botanik#4248 on Discord. Only one submission per druid will be studied.`,
					StartDate: time.Date(2022, time.December, 12, 0, 0, 0, 0, time.UTC),
					EndDate:   time.Date(2023, time.January, 1, 23, 59, 59, 0, time.UTC),
				},
			},
		},
		{
			Number:      2,
			Name:        "Imbolc",
			Description: "The second phase is focused on testing Druids' performance and uptime. Maintenance tasks and upgrades will be performed to test different kinds of state migrations.",
			StartDate:   time.Date(2022, time.January, 2, 0, 0, 0, 0, time.UTC),
			EndDate:     time.Date(2023, time.January, 31, 23, 59, 59, 0, time.UTC),
		},
		{
			Number:      3,
			Name:        "Beltaine",
			Description: "The third phase is all about token dynamics! Druids will engage in various node and community tasks with their precious tokens. Challenges will include some IBC-related tasks to open Nemeton to the interchain world...ime. Maintenance tasks and upgrades will be performed to test different kinds of state migrations.",
			StartDate:   time.Date(2023, time.February, 1, 0, 0, 0, 0, time.UTC),
			EndDate:     time.Date(2023, time.February, 19, 23, 59, 59, 0, time.UTC),
		},
	}
}
