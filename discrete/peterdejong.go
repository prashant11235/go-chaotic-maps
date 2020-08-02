package discrete

import (
	"math"
	"gochaoticmaps/models"
)

type PeterDeJong struct {
	a float64
	b float64
	c float64
	d float64

	init models.Point
	dt float64
	numSteps int
}

func (pdj PeterDeJong) GenerateMapPoints() ([]models.Point) {
	i := 0 
	currPt := pdj.init
	ptArr := make([]models.Point, pdj.numSteps)
	for i < pdj.numSteps {
		pt := pdj.calcNextPoint(currPt)
		ptArr[i] = pt
		currPt = pt 
		i++
	} 

	return ptArr

}

func (pdj PeterDeJong) calcNextPoint(currPt models.Point) (models.Point) {
	xnext := math.Sin(pdj.a * currPt.Y) - math.Cos(pdj.b * currPt.X)
	ynext := math.Sin(pdj.c * currPt.X) - math.Cos(pdj.d * currPt.Y)
	
	pt := models.Point {
		X: xnext, 
		Y: ynext,
		T: currPt.T + 1,
	}
	return pt 
}

func NewPeterDeJong() *PeterDeJong {
	return &PeterDeJong {
		a: -2.24,
		b: 0.43,
		c: -0.65,
		d: -2.43,
		init: models.Point {
			X: 0.1,
			Y: 0.1,
		},
		dt: 1,
		numSteps: 1000,
	}
}
