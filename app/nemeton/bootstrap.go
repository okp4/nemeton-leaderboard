package nemeton

import "time"

func bootstrapPhases() []Phase {
	return []Phase{
		{
			Number:      1,
			Name:        "Sidh",
			Description: "First phase",
			StartDate:   time.Date(2022, time.December, 2, 0, 0, 0, 0, time.UTC),
			EndDate:     time.Date(2023, time.January, 7, 0, 0, 0, 0, time.UTC),
			Tasks: []Task{
				{
					ID:          "0",
					Name:        "Create GenTX",
					Description: "Create your GenTX",
					StartDate:   time.Date(2022, time.December, 2, 0, 0, 0, 0, time.UTC),
					EndDate:     time.Date(2022, time.December, 16, 0, 0, 0, 0, time.UTC),
				},
			},
		},
	}
}
