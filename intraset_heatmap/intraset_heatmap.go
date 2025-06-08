package intraset_heatmap

import "gonum.org/v1/plot/vg"

//intraset heatmap displays color intensity for weight
//n'th rep is displayed on the y axis
//sets are shown as seperate columns stacked next to each other for different x values
//the x axis is the session date

//x is date
//y is rep number
//color is weight lifted

type Sessioner interface {
	GetSession(i int) *Session // GetSessions retrieves all sessions
	Len() int                 // Len returns the number of sessions
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
	Sessions
	 
	ColumnWidth, MaxHeight, MinHeight vg.Length // Max and Min height for the vectoral height of each column
	MaxReps, MinReps     int       // Max and Min reps across all sessions

}

func NewIntrasetHeatmap( data Sessioner, columnWidth,minHeight, maxHeight vg.Length) *IntrasetHeatmap {
	cpy := CopySessions(data)
	maxReps := 0
	minReps := 0
	for _, session := range cpy { // Iterate through each session to find max and min reps
		if len(session.Sets) == 0 {
			continue // Skip empty sessions
		}
		for _, set := range session.Sets {
			if len(set.Reps) == 0 {
				continue // Skip empty sets
			}
			maxReps = max(maxReps, len(set.Reps)) // Update max reps
			minReps = min(minReps, len(set.Reps)) // Update min reps
		}

	}
	return &IntrasetHeatmap{
		Sessions:  cpy,
		MaxHeight: maxHeight,
		MinHeight: minHeight,
		MaxReps:   maxReps,
		MinReps:   minReps,
		ColumnWidth: columnWidth,
	}
}
