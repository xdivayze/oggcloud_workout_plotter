package intraset_heatmap

type Set struct {
	Reps []*Rep  `json:"reps"` // the reps in this set
	SetNo int    `json:"set_no"` // the set number in the session, starting from 1
	
}

