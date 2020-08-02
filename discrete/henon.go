package discrete

import (
	"gochaoticmaps/models"
)

type Henon struct {
	a float64
	b float64
	init models.Point
	dt float64
	numSteps int
}

func (h Henon) GenerateMapPoints() ([]models.Point) {
	i := 0 
	currPt := h.init
	ptArr := make([]models.Point, h.numSteps)
	for i < h.numSteps {
		pt := h.calcNextPoint(currPt)
		ptArr[i] = pt
		currPt = pt 
		i++
	} 

	return ptArr

}

func (h Henon) calcNextPoint(currPt models.Point) (models.Point) {
	xnext := 1 - h.a*currPt.X*currPt.X + currPt.Y
	ynext := h.b * currPt.X
	
	pt := models.Point {
		X: xnext, 
		Y: ynext,
		T: currPt.T + 1,
	}
	return pt 
}

func NewHenon() *Henon {
	return &Henon {
		a: 1.4,
		b: 0.3,
		init: models.Point {
			X: 1,
			Y: 1,
		},
		dt: 1,
		numSteps: 1000,
	}
}
