package discrete

import (
	"gochaoticmaps/models"
)

type Duffing struct {
	a float64
	b float64
	init models.Point 
	numSteps int
}

// Initial values
var x = 0.1
var y = 0.1

// Time variation
var dt = 0.04

func (d Duffing) GenerateMapPoints() ([]models.Point) {
	i := 0 
	ptArr := make([]models.Point, d.numSteps)
	for i < d.numSteps {
		pt := d.calcNextPoint()
		ptArr[i] = pt
		i++
	} 

	return ptArr

}

func (d Duffing) calcNextPoint() (models.Point) {
	xnext := y
	ynext := -d.b*x + d.a*y - y*y*y
	
	pt := models.Point {
		X: xnext, 
		Y: ynext,
	}

	x = xnext
	y = ynext 

	return pt 
}

func NewDuffing() *Duffing {
	x = 0.1
	y = 0.1

	return &Duffing{
		a: 2.75,
		b: 0.2,
		init: models.Point {
			X: 0.1,
			Y: 0.1,
		},
		numSteps: 1000,
	}
}
