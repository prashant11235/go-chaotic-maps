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

// Initial values
var x float64
var y float64
var z float64

func (l Lorenz) GenerateMapPoints() ([]models.Point) {
	i := 0 
	ptArr := make([]models.Point, l.numSteps)
	for i < l.numSteps {
		pt := l.calcNextPoint()
		ptArr[i] = pt
		i++
	} 

	return ptArr 
}

func (l Lorenz) calcNextPoint() (models.Point) {
	
	dx := l.sigma * (y - x)
	dy := x * (l.rho -z) - y
	dz := x * y - l.beta * z

	x += dx * l.dt
	y += dy * l.dt
	z += dz * l.dt

	pt := models.Point {
		X: x, 
		Y: y,
		Z: z,
	}

	return pt 
}

func NewLorenz() *Lorenz {
	x = 0.1
	y = 0.1
	z = 0.1

	return &Lorenz{
		sigma: 10.0,
		rho: 28.0,
		beta: 8.0/3.0,
		init: models.Point {
			X: 0.1,
			Y: 0.1,
			Z: 0.1,
		},
		dt: 0.01,
		numSteps: 2000,
	}
}
