package nemeton

import "time"

// nolint: funlen,lll
func bootstrapPhases() []Phase {
	r100 := 100
	r200 := 200
	r500 := 500
	return []Phase{
		{
			Number:      1,
			Name:        "Pre-Sidh",
			Description: "Pre phase.",
			StartDate:   time.Date(2022, time.October, 2, 0, 0, 0, 0, time.UTC),
			EndDate:     time.Date(2022, time.November, 7, 0, 0, 0, 0, time.UTC),
			Tasks: []Task{
				{
					ID:          "0",
					Type:        taskTypeBasic,
					Name:        "Say hello",
					Description: "Say hello to the community",
					StartDate:   time.Date(2022, time.October, 2, 0, 0, 0, 0, time.UTC),
					EndDate:     time.Date(2022, time.November, 7, 0, 0, 0, 0, time.UTC),
					Rewards:     &r100,
				},
				{
					ID:          "1",
					Type:        taskTypeBasic,
					Name:        "Say goodbye",
					Description: "Say goodbye to the community",
					StartDate:   time.Date(2022, time.October, 2, 0, 0, 0, 0, time.UTC),
					EndDate:     time.Date(2022, time.November, 7, 0, 0, 0, 0, time.UTC),
					Rewards:     &r100,
				},
				{
					ID:          "2",
					Type:        taskTypeBasic,
					Name:        "Be gentle",
					Description: "Be gentle with the community",
					StartDate:   time.Date(2022, time.October, 2, 0, 0, 0, 0, time.UTC),
					EndDate:     time.Date(2022, time.November, 7, 0, 0, 0, 0, time.UTC),
					Rewards:     &r500,
				},
				{
					ID:          "3",
					Type:        taskTypeUptime,
					Name:        "Uptime",
					Description: "Maintain the best uptime!",
					StartDate:   time.Date(2022, time.October, 2, 0, 0, 0, 0, time.UTC),
					EndDate:     time.Date(2022, time.November, 7, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		{
			Number:      2,
			Name:        "Sidh",
			Description: "This first phase is pretty basic, it is dedicated to setting up Druids' validator environment, participating in the genesis, and getting familiar with the OKP4 testnet.",
			StartDate:   time.Date(2022, time.November, 2, 0, 0, 0, 0, time.UTC),
			EndDate:     time.Date(2022, time.December, 7, 0, 0, 0, 0, time.UTC),
			Tasks: []Task{
				{
					ID:          "0",
					Type:        taskTypeBasic,
					Name:        "Submit your gentx",
					Description: "Submit your gentx in time.",
					StartDate:   time.Date(2022, time.November, 2, 0, 0, 0, 0, time.UTC),
					EndDate:     time.Date(2022, time.November, 15, 0, 0, 0, 0, time.UTC),
					Rewards:     &r200,
				},
				{
					ID:          "1",
					Type:        taskTypeBasic,
					Name:        "Setup your node",
					Description: "Make your node join the network.",
					StartDate:   time.Date(2022, time.November, 15, 0, 0, 0, 0, time.UTC),
					EndDate:     time.Date(2022, time.December, 7, 0, 0, 0, 0, time.UTC),
					Rewards:     &r500,
				},
				{
					ID:          "2",
					Type:        taskTypeTweetNemeton,
					Name:        "Nemeton Tweet",
					Description: "Tweet about Nemeton.",
					StartDate:   time.Date(2022, time.November, 2, 0, 0, 0, 0, time.UTC),
					EndDate:     time.Date(2022, time.December, 7, 0, 0, 0, 0, time.UTC),
					Rewards:     &r100,
				},
				{
					ID:          "3",
					Type:        taskTypeUptime,
					Name:        "Uptime",
					Description: "Maintain the best uptime",
					StartDate:   time.Date(2022, time.November, 2, 0, 0, 0, 0, time.UTC),
					EndDate:     time.Date(2022, time.December, 7, 0, 0, 0, 0, time.UTC),
				},
				{
					ID:          "4",
					Type:        taskTypeSubmission,
					Name:        "Documentation enhancement",
					Description: "Help us enhance the node documentation.",
					StartDate:   time.Date(2022, time.November, 2, 0, 0, 0, 0, time.UTC),
					EndDate:     time.Date(2022, time.December, 7, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		{
			Number:      3,
			Name:        "Imbolc",
			Description: "The second phase is focused on testing Druids' performance and uptime. Maintenance tasks and upgrades will be performed to test different kinds of state migrations.",
			StartDate:   time.Date(2022, time.December, 2, 0, 0, 0, 0, time.UTC),
			EndDate:     time.Date(2023, time.January, 7, 0, 0, 0, 0, time.UTC),
		},
		{
			Number:      4,
			Name:        "Beltaine",
			Description: "The third phase is all about token dynamics! Druids will engage in various node and community tasks with their precious tokens. Challenges will include some IBC-related tasks to open Nemeton to the interchain world...ime. Maintenance tasks and upgrades will be performed to test different kinds of state migrations.",
			StartDate:   time.Date(2023, time.January, 2, 0, 0, 0, 0, time.UTC),
			EndDate:     time.Date(2023, time.February, 7, 0, 0, 0, 0, time.UTC),
		},
	}
}
