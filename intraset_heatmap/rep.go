package intraset_heatmap

import "image/color"

type Rep struct {
	Weight float64 `json:"weight"` // weight lifted in this rep
	RepNo int     `json:"rep_no"` // the rep number in the set, starting from 1
	Color color.Color  `json:"color"` // color intensity for the weight lifted, used for heatmap visualization
}

func GenerateColor(weight float64) color.Color {
	// This function generates a color based on the weight lifted.
	
	intensity := uint8(255 - ((weight*255)/120) ) // Assuming max weight is 120 for full intensity
	return color.RGBA{intensity, 0, 0, 255} 
}

func NewRep(weight float64, repNo int) *Rep {
	return &Rep{
		Weight: weight,
		RepNo:  repNo,
		Color:  GenerateColor(weight), // Generate color based on weight
	}
}