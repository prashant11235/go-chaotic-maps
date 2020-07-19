package continuous

import (
	"math"
	"gochaoticmaps/models"
	"gochaoticmaps/vis"
)

// Constants
type DuffingVisContext struct {
	vp vis.VisualizeParams
}

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

func GetDuffingVisualizeContext() vis.VisualizeContext {
	lorenzVis := DuffingVisContext{}
	lorenzVis.vp = vis.VisualizeParams {
			SizeX: 750,
			SizeY: 750,
			IgnoreAxis: "Z",
	}
	
	return lorenzVis
}

func (l DuffingVisContext) ScaleX(x float64) float64 {
	return l.InitX(l.vp.SizeX) + 200 * x;
}

func (l DuffingVisContext) ScaleY(y float64) float64 {
	return l.InitY(l.vp.SizeX) + 200 * y;
}

func (l DuffingVisContext) ScaleZ(z float64) float64 {
	return l.InitZ(l.vp.SizeX) + 10 * z;
}

func (l DuffingVisContext) VisualizeParams() vis.VisualizeParams {
	return l.vp
}

func (l DuffingVisContext) InitX(x float64) float64 {
	return l.vp.SizeX/2
}

func (l DuffingVisContext) InitY(y float64) float64 {
	return l.vp.SizeX/2
}

func (l DuffingVisContext) InitZ(z float64) float64 {
	return l.vp.SizeX/2
}

func (l DuffingVisContext) ScaleFactor(sizeX float64) float64 {
	return sizeX/55
}