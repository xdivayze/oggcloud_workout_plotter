package intraset_heatmap

import (
	"fmt"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

func (i *IntrasetHeatmap) Plot(c draw.Canvas, plt *plot.Plot) { //TODO dynamically calculated col width given a max session col size

	trX, _ := plt.Transforms(&c)

	//calculate the total width of the plot
	nSession := i.Len()

	totalWidthWithoutSessionSpacing := c.Size().X
	fmt.Printf("canvas size is %f, total width without session spacing is %f\n", c.Size().X, totalWidthWithoutSessionSpacing)

	perSessionWidth := (totalWidthWithoutSessionSpacing - vg.Length((nSession-1)*4)) / vg.Length(nSession)

	for _, session := range i.Sessions {
		setCount := len(session.Sets)

		sessionSetColWidth := perSessionWidth / vg.Length(setCount)
		sessionDateUnix := float64(session.Date.Unix())
		//margin := ((totalWidth/ vg.Length(nSession)) - sessionSetColWidth)/2

		sessionDateString := session.Date.Format("2006-01-02")
		fmt.Printf("Plotting session on %s \n", sessionDateString)
		colx := trX(sessionDateUnix)
		offset := sessionSetColWidth * vg.Length(setCount) / 2
		x := colx - offset // bottom left corner of the first set in the session
		for _, set := range session.Sets {
			fmt.Printf("Plotting set %d for session on %s at x: %f\n", set.SetNo, session.Date, x)

			for _, rep := range set.Reps {

				rect := vg.Rectangle{
					Min: vg.Point{
						X: x,
						Y: i.height(rep.RepNo), // Height for each rep
					},
					Max: vg.Point{
						X: x + sessionSetColWidth,
						Y: i.height(rep.RepNo + 1), // Height for each rep
					}}
				p := rect.Path()
				c.SetColor(rep.Color) // Set the color for the rectangle
				p.Close()
				c.Fill(p) // Fill the rectangle with the color

			}
			x += sessionSetColWidth // Move to the next set's x position

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
