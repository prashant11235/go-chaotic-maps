package continuous

import (
	"gochaoticmaps/models"
	"gochaoticmaps/vis"
)

type VanderpolVisContext struct {
	vp vis.VisualizeParams
}

type Vanderpol struct {
	b float64
	
	init models.Point
	dt float64
	numSteps int
}
