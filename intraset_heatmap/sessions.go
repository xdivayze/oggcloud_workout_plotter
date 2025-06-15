package intraset_heatmap

import "time"

type Sessioner interface {
	GetSession(i int) *Session // GetSessions retrieves i. session
	Len() int                  // Len returns the number of sessions
	GetMaxSetSize() int      // GetMaxSetSize returns the maximum number of sets in any session
	GetDateRange() (min, max time.Time) // GetDateRange returns the minimum and maximum dates from the sessions
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

func (s Sessions) GetMaxSetSize() int {
	if len(s) == 0 {
		return 0 // Return 0 if there are no sessions
	}

	maxSetSize := 0
	for _, session := range s {
		maxSetSize = max(maxSetSize, len(session.Sets))
	}
	return maxSetSize
}

// GetDateRange returns the minimum and maximum dates from the sessions, truncated to the start of the day.
// If there are no sessions, it returns zero values for both dates.
func (s Sessions) GetDateRange() (min, max time.Time) {
	if len(s) == 0 {
		return time.Time{}, time.Time{} // Return zero values if no sessions
	}

	min = s[0].Date
	max = s[0].Date

	for _, session := range s {
		if session.Date.Before(min) {
			min = session.Date
		}
		if session.Date.After(max) {
			max = session.Date
		}
	}
	return min.Truncate(24 * time.Hour), max.Truncate(24 * time.Hour) // Truncate to the start of the day
}

func CopySessions(data Sessioner) Sessions { //TODO make dates in increasing order
	copied := make(Sessions, data.Len())
	for i := 0; i < data.Len(); i++ {
		copied[i] = data.GetSession(i)
	}
	return copied
}
