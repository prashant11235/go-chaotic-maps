package vis

import (
	"gochaoticmaps/models"
)

type Visualizer interface {
	Visualize(cm models.ChaoticMap, dv DefaultVisualization)
}

type VisualizeParams struct {
	Size models.MapSize
	IgnoreAxis string // for a 3d point one axis needs to be ingnored in plot
	PlotType string // path or circle etc
	// Keep adding all stuff needed to visualize a group of points in space
}

type VisualizeContext interface {
	InitPt(models.MapSize) models.Point
	VisualizeParams() VisualizeParams
}

type DefaultVisualization struct {
	vp VisualizeParams
}

func (dv DefaultVisualization) VisualizeParams() VisualizeParams {
	return dv.vp
}

func (dv DefaultVisualization) InitPt(mapSize models.MapSize) models.Point {
	return models.Point {
		X: float64(mapSize.SizeX)/2,
		Y: float64(mapSize.SizeY)/2,
		Z: float64(mapSize.SizeZ)/2,
		T: 0,
	}
}

func GetDefaultVisualizeContext() VisualizeContext {
	defaultVis := DefaultVisualization{}
	defaultVis.vp = VisualizeParams {
		Size: models.MapSize{
			SizeX: 500,
			SizeY: 250,
			SizeZ: 250,				
		},
		IgnoreAxis: "Y",
	}
	
	return defaultVis
}
