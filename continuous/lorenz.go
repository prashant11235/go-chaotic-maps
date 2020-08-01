package continuous

import (
	"gochaoticmaps/models"
)

type Lorenz struct {
	sigma float64
	rho float64
	beta float64

	init models.Point
	dt float64
	numSteps int
}

func (l Lorenz) GenerateMapPoints() ([]models.Point) {
	i := 0 
	currPt := l.init
	ptArr := make([]models.Point, l.numSteps)
	for i < l.numSteps {
		pt := l.calcNextPoint(currPt)
		ptArr[i] = pt
		currPt = pt 
		i++
	} 

	return ptArr 
}

func (l Lorenz) calcNextPoint(currPt models.Point) (models.Point) {
	
	dx := l.sigma * (currPt.Y - currPt.X)
	dy := currPt.X * (l.rho - currPt.Z) - currPt.Y
	dz := currPt.X * currPt.Y - l.beta * currPt.Z

	x := currPt.X + dx * l.dt
	y := currPt.Y + dy * l.dt
	z := currPt.Z + dz * l.dt
	t := currPt.T + l.dt 

	pt := models.Point {
		X: x, 
		Y: y,
		Z: z,
		T: t,
	}

	return pt 
}

func NewLorenz() *Lorenz {
	return &Lorenz{
		sigma: 10.0,
		rho: 28.0,
		beta: 8.0/3.0,
		init: models.Point {
			X: 10,
			Y: 10,
			Z: 10,
			T: 0.0,
		},
		dt: 0.01,
		numSteps: 10000,
	}
}
