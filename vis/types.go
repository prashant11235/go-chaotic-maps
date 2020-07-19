package vis

import (
	"gochaoticmaps/models"
)

type Visualizer interface {
	Visualize(cm models.ChaoticMap, vc VisualizeContext)
}

type VisualizeParams struct {
	SizeX float64 
	SizeY float64
	IgnoreAxis string // for a 3d point one axis needs to be ingnored in plot
	PlotType string // path or circle etc
	// Keep adding all stuff needed to visualize a group of points in space
}

type VisualizeContext interface {
	ScaleX(float64) float64
	ScaleY(float64) float64
	ScaleZ(float64) float64
	ScaleFactor(float64) float64
	InitX(float64) float64
	InitY(float64) float64
	InitZ(float64) float64
	VisualizeParams() VisualizeParams
}
