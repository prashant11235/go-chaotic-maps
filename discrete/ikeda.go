package discrete

import (
	"gochaoticmaps/models"
)

type Ikeda struct {
	a float64
	b float64
	init models.Point
	dt float64
	numSteps int
}
