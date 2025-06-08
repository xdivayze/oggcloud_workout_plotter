package intraset_heatmap

import (
	"fmt"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

func (i *IntrasetHeatmap) Plot(c draw.Canvas, plt *plot.Plot) { //TODO dynamically calculated col width given a max session col size

	trX, _ := plt.Transforms(&c)
	for _, session := range i.Sessions { 
		sessionDateUnix := float64(session.Date.Unix())
		setCount := len(session.Sets)

		colx := trX(sessionDateUnix)
		offset := i.ColumnWidth * vg.Length(setCount) / 2
		x := colx - offset // Center the first set of the session
		for _, set := range session.Sets {
			x += i.ColumnWidth // Move to the next set's x position
			fmt.Printf("Plotting set %d for session on %s at x: %f\n", set.SetNo, session.Date, x)

			for _, rep := range set.Reps {

				rect := vg.Rectangle{
					Min: vg.Point{
						X: x,
						Y: i.height(rep.RepNo), // Height for each rep
					},
					Max: vg.Point{
						X: x + i.ColumnWidth,
						Y: i.height(rep.RepNo + 1), // Height for each rep
					}}
				p := rect.Path()
				c.SetColor(rep.Color) // Set the color for the rectangle
				p.Close()
				c.Fill(p) // Fill the rectangle with the color

			}

		}
	}
}

// basic linear interpolation
func (i *IntrasetHeatmap) height(repCount int) vg.Length {
	minRep := vg.Length(i.MinReps)
	maxRep := vg.Length(i.MaxReps)

	slope := (i.MaxHeight - i.MinHeight) / (maxRep - minRep)
	yIntercept := i.MaxHeight - slope*maxRep //calculate y intercept to write linear formula
	return vg.Length(repCount)*slope + yIntercept
}
