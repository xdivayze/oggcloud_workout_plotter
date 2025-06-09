package intraset_heatmap

import (
	"math"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
)

//intraset heatmap displays color intensity for weight
//n'th rep is displayed on the y axis
//sets are shown as seperate columns stacked next to each other for different x values
//the x axis is the session date

//x is date
//y is rep number
//color is weight lifted

type Sessioner interface {
	GetSession(i int) *Session // GetSessions retrieves all sessions
	Len() int                  // Len returns the number of sessions
}

type Sessions []*Session

func (s Sessions) GetSession(i int) *Session {
	if i < 0 || i >= len(s) {
		return &Session{} // Return an empty session if index is out of bounds
	}
	return s[i]
}

func (s Sessions) Len() int {
	return len(s)
}

func CopySessions(data Sessioner) Sessions { //TODO make dates in increasing order
	copied := make(Sessions, data.Len())
	for i := 0; i < data.Len(); i++ {
		copied[i] = data.GetSession(i)
	}
	return copied
}

type IntrasetHeatmap struct {
	TotalSetCount int // Total number of sets across all sessions
	Sessions

	ColumnWidth, MaxHeight, MinHeight vg.Length // Max and Min height for the vectoral height of each column
	MaxReps, MinReps                  int       // Max and Min reps across all sessions

}



func (i *IntrasetHeatmap) GenerateXTickers(min, max float64) []plot.Tick {
	const layout = "Jan 02"
	var ticks []plot.Tick
	const interval = float64(24 * time.Hour / time.Second)
	for t := math.Ceil(min/interval) * interval; t <= max; t += interval {
		ticks = append(ticks, plot.Tick{
			Value: t,
			Label: time.Unix(int64(t), 0).Format(layout),
		})
	}
	return ticks
}

func NewIntrasetHeatmap(data Sessioner, columnWidth, minHeight, maxHeight vg.Length) *IntrasetHeatmap {
	cpy := CopySessions(data)
	maxReps := 0
	minReps := 0
	TotalSetCount := 0
	for _, session := range cpy { // Iterate through each session to find max and min reps
		if len(session.Sets) == 0 {
			continue // Skip empty sessions
		}
		TotalSetCount += len(session.Sets) // Count total sets
		for _, set := range session.Sets {
			if len(set.Reps) == 0 {
				continue // Skip empty sets
			}
			maxReps = max(maxReps, len(set.Reps)) // Update max reps
			minReps = min(minReps, len(set.Reps)) // Update min reps
		}
		session.Date = session.Date.Truncate(24 * time.Hour) // Normalize date to midnight

	}
	return &IntrasetHeatmap{
		Sessions:      cpy,
		MaxHeight:     maxHeight,
		MinHeight:     minHeight,
		MaxReps:       maxReps,
		MinReps:       minReps,
		ColumnWidth:   columnWidth,
		TotalSetCount: TotalSetCount,
	}
}
