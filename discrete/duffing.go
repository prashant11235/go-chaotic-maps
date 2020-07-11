package discrete 

import (
	"math"
	"gochaoticmaps/models"
)

// Constants

type Duffing struct {
	alpha float64
	beta float64
	delta float64
	gamma float64
	omega float64
	numSteps int
}

// Initial values
var x = 0.0
var y = 0.0
var t = 0.0

// Time variation
var dt = 0.04

func GeneratePoints() ([]models.Point) {
	i := 0 
	ptArr := make([]Point, maxSteps)
	for i < maxSteps {
		pt := calcNextPoint()
		ptArr[i] = pt
		i++
	} 

	return ptArr 

}

func calcNextPoint() (models.Point) {
	dx := float64(y)
	dy := -delta*float64(y) - beta*float64(x) - alpha*float64(x*x*x) + gamma*(float64(math.Cos(omega*t)))
	
	x += dx * dt
	y += dy * dt
	
	t += dt
	
	pt := models.Point {
		X: x, 
		Y: y,
	}

	return pt 
}

func NewDuffing() *Duffing {
	x = 0.1
	y = 0.1
	z = 0.1

	return &Duffing{
		alpha: 10.0,
		beta: 28.0,
		delta: 8.0/3.0,
		gamma: 1,
		omega: 2,
		init: models.Point {
			X: 0.1,
			Y: 0.1,
			Z: 0.1,
		},
		dt: 0.01,
		numSteps: 10000,
	}
}
