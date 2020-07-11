package draw 

import (
	"fmt"
	"gochaoticmaps/models"
)

type Draw struct {
	SizeX int 
	SizeY int 
	InitX int 
	InitY int 
}

func (d Draw) GenPath(pts []models.Point) (string) {
	path := "M " + fmt.Sprintf("%f", d.scaleX(pts[0].X)) + "," + fmt.Sprintf("%f", d.scaleZ(pts[0].Z)) + " "
	for _, pt := range pts {
		currX := d.scaleX(pt.X)
		currZ := d.scaleZ(pt.Z)
		
		path +=  " L " + fmt.Sprintf("%f", currX) + "," + fmt.Sprintf("%f", currZ)
	}
	return path 
}

func (d Draw) scaleX(x float64) float64 {
	return 450 + 20 * x;
}

func (d Draw) scaleZ(z float64) float64 {
	return 50 + 20 * z;
}

func NewDraw() *Draw {
	return &Draw {
		SizeX: 1000,
		SizeY: 1000,
		InitX: 450,
		InitY: 50,
	}
}
