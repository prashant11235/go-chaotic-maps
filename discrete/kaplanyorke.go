package discrete

import (
	"math"
	"gochaoticmaps/models"
)

type KaplanYorke struct {
	alpha float64
	b float64
	init models.Point
	dt float64
	numSteps int
}

func (ky KaplanYorke) GenerateMapPoints() ([]models.Point) {
	i := 0 
	currPt := ky.init
	ptArr := make([]models.Point, ky.numSteps)
	for i < ky.numSteps {
		pt := ky.calcNextPoint(currPt)
		ptArr[i] = pt
		currPt = pt 
		i++
	} 

	return ptArr

}

func (ky KaplanYorke) calcNextPoint(currPt models.Point) (models.Point) {
	xnext := math.Mod(2*currPt.X, 0.99995)
	ynext := ky.alpha*currPt.Y + math.Cos(4*math.Pi*currPt.X)
	
	pt := models.Point {
		X: xnext, 
		Y: ynext,
		T: currPt.T + 1,
	}
	return pt 
}

func NewKaplanYorke() *KaplanYorke {
	return &KaplanYorke {
		alpha: 0.02,
		init: models.Point {
			X: 0.367812,
			Y: 0.667751,
		},
		dt: 1,
		numSteps: 1000,
	}
}
