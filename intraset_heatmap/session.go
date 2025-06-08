package intraset_heatmap

import "time"

type Session struct {
	Sets []*Set `json:"sets"` // the sets in this session
	Date time.Time `json:"date"` // the date of the session in YYYY-MM-DD format
	Volume float64 `json:"total_volume"` // volume lifted in this session
}

func NewSession(sets []*Set, time time.Time) *Session {
	volume := 0.0
	for _, set := range sets {
		for _, rep := range set.Reps {
			volume += rep.Weight
		}
	}

	return &Session{
		Sets: sets,
		Date: time, // Default to current time, can be set later
		Volume: volume, // Default volume, can be calculated later
	}
}