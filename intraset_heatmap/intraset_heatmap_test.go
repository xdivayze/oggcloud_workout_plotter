package intraset_heatmap_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/xdivayze/oggcloud_workout_plotter/intraset_heatmap"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
)

func TestIntrasetHeatmap(t *testing.T) {
	require := require.New(t)

	rand := rand.New(rand.NewSource(0))
	data := intraset_heatmap.Sessions(generateRandomData(5, 70, 20, 11, 5, 2, 5, rand))

	p := plot.New()

	p.Title.Text = "Intraset Heatmap Test"
	p.X.Label.Text = "Session Date"
	p.Y.Label.Text = "Rep Number"

	heatmap := intraset_heatmap.NewIntrasetHeatmap(data, 200, 50, vg.Points(10), vg.Points(50), vg.Points(200))

	p.Add(heatmap)
	start := time.Now().AddDate(0, 0, -4).Truncate(24 * time.Hour) // Start 4 days ago, at midnight
	end := time.Now().AddDate(0, 0, 1).Truncate(24 * time.Hour)    // End today, at midnight

	p.Y.Padding = vg.Length(40)

	p.X.Tick.Marker = plot.TickerFunc(heatmap.GenerateXTickers)

	p.X.Max = float64(end.Unix())   // Set max x to 5 days ago
	p.X.Min = float64(start.Unix()) // Set min x to 5 days ago

	p.Y.Max = float64(heatmap.MaxReps + 5)
	p.Y.Min = 1

	require.Nil(p.Save(10*vg.Centimeter, 10*vg.Centimeter, "intraset_heatmap_test.png"))

}

func generateRandomData(n int, maxWeight, minWeight float64, maxReps, minReps int,
	minSetCount, maxSetCount int, rand *rand.Rand) []*intraset_heatmap.Session {
	data := make([]*intraset_heatmap.Session, n)
	for i := 0; i < n; i++ {
		data[i] = intraset_heatmap.NewSession(generateRandomSets(
			rand.Intn(maxSetCount-minSetCount+1)+minSetCount,
			maxWeight, minWeight, maxReps, minReps), time.Now().AddDate(0, 0, -i)) // Set date to today minus i days

	}
	return data

}

func generateRandomSets(n int, maxWeight, minWeight float64, maxRepCount, minRepCount int) []*intraset_heatmap.Set {
	sets := make([]*intraset_heatmap.Set, n)
	for i := 0; i < n; i++ {
		setNo := i + 1
		repsCount := rand.Intn(maxRepCount-minRepCount) + minRepCount
		reps := generateRandomReps(repsCount, maxWeight, minWeight)
		sets[i] = &intraset_heatmap.Set{
			Reps:  reps,
			SetNo: setNo,
		}
	}
	return sets
}

func generateRandomReps(n int, maxWeight float64, minWeight float64) []*intraset_heatmap.Rep {

	reps := make([]*intraset_heatmap.Rep, n)
	for i := 0; i < n; i++ {
		repNo := i
		rep := intraset_heatmap.NewRep((minWeight + rand.Float64()*(maxWeight-minWeight)), repNo)
		reps[i] = rep
	}
	return reps
}
