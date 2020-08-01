package continuous

import (
	"math"
	"gochaoticmaps/models"
)

type Thomas struct {
	b float64
	
	init models.Point
	numParticles int
	dt float64
	numSteps int
}

func (th Thomas) GenerateMapPoints() ([]models.Point) {
	i := 0
	j := 1
	ptArr := make([]models.Point, th.numSteps * th.numParticles)
	currPt := th.init
	for j <= th.numParticles {
		currPt.X = float64(j) * 0.1
		for i < th.numSteps {
			pt := th.calcNextPoint(currPt)
			ptArr[i] = pt
			currPt = pt 
			i++
		} 
		j++
	}
	return ptArr 
}

func (th Thomas) calcNextPoint(currPt models.Point) (models.Point) {

	dx := math.Sin(currPt.Y) - th.b * currPt.X
    dy := math.Sin(currPt.Z) - th.b * currPt.Y
    dz := math.Sin(currPt.X) - th.b * currPt.Z
    
	x := currPt.X + dx * th.dt
	y := currPt.Y + dy * th.dt
	z := currPt.Z + dz * th.dt
	t := currPt.T + th.dt 

	pt := models.Point {
		X: x, 
		Y: y,
		Z: z,
		T: t,
	}

	return pt 
}

func NewThomas() *Thomas {
	return &Thomas{
		b: 0.18,
		numParticles: 10,
		init: models.Point {
			X: 0.1,
			Y: 0.0,
			Z: 0.0,
			T: 0.0,
		},
		dt: 0.1,
		numSteps: 10000,
	}
}
