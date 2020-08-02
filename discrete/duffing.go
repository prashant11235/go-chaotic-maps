package discrete

import (
	"gochaoticmaps/models"
)

type Duffing struct {
	a float64
	b float64
	init models.Point 
	dt float64
	numSteps int
}

func (d Duffing) GenerateMapPoints() ([]models.Point) {
	i := 0 
	currPt := d.init
	ptArr := make([]models.Point, d.numSteps)
	for i < d.numSteps {
		pt := d.calcNextPoint(currPt)
		ptArr[i] = pt
		currPt = pt 
		i++
	} 

	return ptArr

}

func (d Duffing) calcNextPoint(currPt models.Point) (models.Point) {
	xnext := currPt.Y
	ynext := -d.b*currPt.X + d.a*currPt.Y - currPt.Y*currPt.Y*currPt.Y
	
	pt := models.Point {
		X: xnext, 
		Y: ynext,
	}
	return pt 
}

func NewDuffing() *Duffing {
	return &Duffing{
		a: 2.75,
		b: 0.2,
		init: models.Point {
			X: 0.1,
			Y: 0.1,
		},
		dt: 0.01,
		numSteps: 1000,
	}
}
