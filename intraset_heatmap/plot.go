package intraset_heatmap

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

// TODO calculate the max and min x's of the designated are so that when excess
// TODO space is available(min x is not exactly the min date) the column groups do not overlap
func (i *IntrasetHeatmap) Plot(c draw.Canvas, plt *plot.Plot) {

	trX, trY := plt.Transforms(&c)
	totalSize := c.Rectangle.Size().X
	i.ColumnWidth = totalSize / vg.Length(i.divisor)

	for _, session := range i.Sessions {
		setCount := len(session.Sets)

		sessionDateUnix := float64(session.Date.Unix())

		colx := trX(sessionDateUnix)
		offset := i.ColumnWidth * vg.Length(setCount) / 2
		x := colx - offset // bottom left corner of the first set in the session
		for _, set := range session.Sets {

			for _, rep := range set.Reps {

				rect := vg.Rectangle{
					Min: vg.Point{
						X: x,
						Y: trY(float64(rep.RepNo)), // Height for each rep
					},
					Max: vg.Point{
						X: x + i.ColumnWidth,
						Y: trY(float64(rep.RepNo + 1)), // Height for each rep
					}}
				p := rect.Path()
				c.SetColor(rep.Color) // Set the color for the rectangle
				p.Close()
				c.Fill(p) // Fill the rectangle with the color

			}
			x += i.ColumnWidth // Move to the next set's x position

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
