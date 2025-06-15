package intraset_heatmap

import "image/color"

type Rep struct {
	Weight float64     `json:"weight"` // weight lifted in this rep
	RepNo  int         `json:"rep_no"` // the rep number in the set, starting from 1
	Color  color.Color `json:"color"`  // color intensity for the weight lifted, used for heatmap visualization
}

func GenerateColor(colorInterpolation func(weight float64) color.Color, weight float64) color.Color {
	// This function generates a color based on the weight lifted.
	return colorInterpolation(weight)
}

func NewRep(weight float64, repNo int) *Rep { //color should be generated based on weight later using heatmap's method
	return &Rep{
		Weight: weight,
		RepNo:  repNo,
	}
}
