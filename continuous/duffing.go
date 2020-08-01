package continuous

import (
	"math"
	"gochaoticmaps/models"
)

type Duffing struct {
	alpha float64
	beta float64
	delta float64
	gamma float64
	omega float64
	init models.Point
	numSteps int
	dt float64
	currTime float64
}

func (d Duffing) GenerateMapPoints() ([]models.Point) {
	i := 0 
	currPt := d.init 
	currTime := d.currTime
	ptArr := make([]models.Point, d.numSteps)
	for i < d.numSteps {
		pt := d.calcNextPoint(currPt, currTime)
		ptArr[i] = pt
		currPt = pt 
		currTime += d.dt
		i++
	} 

	return ptArr 

}

func (d Duffing) calcNextPoint(currPt models.Point, currTime float64) (models.Point) {
	dx := float64(currPt.Y)
	dy := -d.delta*float64(currPt.Y) - d.beta*float64(currPt.X) - d.alpha*float64(currPt.X*currPt.X*currPt.X) + d.gamma*(float64(math.Cos(d.omega*currTime)))
	
	x := currPt.X + dx * d.dt
	y := currPt.Y + dy * d.dt
	
	pt := models.Point {
		X: x, 
		Y: y,
	}

	return pt 
}

func NewDuffing() *Duffing {
	return &Duffing{
		alpha: 1.0,
		beta: -1.0,
		delta: 0.2,
		gamma: 0.3,
		omega: 1.0,
		init: models.Point {
			X: 0.0,
			Y: 0.0,
			Z: 0.0,
		},
		dt: 0.04,
		currTime: 0,
		numSteps: 50000,
	}
}
