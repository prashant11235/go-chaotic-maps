package continuous

import (
	"gochaoticmaps/models"
	"gochaoticmaps/vis"
)

type LorenzVisContext struct {
	vp vis.VisualizeParams
}

type Lorenz struct {
	sigma float64
	rho float64
	beta float64

	init models.Point
	dt float64
	numSteps int
}

func (l Lorenz) GenerateMapPoints() ([]models.Point) {
	i := 0 
	currPt := l.init
	ptArr := make([]models.Point, l.numSteps)
	for i < l.numSteps {
		pt := l.calcNextPoint(currPt)
		ptArr[i] = pt
		currPt = pt 
		i++
	} 

	return ptArr 
}

func (l Lorenz) calcNextPoint(currPt models.Point) (models.Point) {
	
	dx := l.sigma * (currPt.Y - currPt.X)
	dy := currPt.X * (l.rho - currPt.Z) - currPt.Y
	dz := currPt.X * currPt.Y - l.beta * currPt.Z

	x := currPt.X + dx * l.dt
	y := currPt.Y + dy * l.dt
	z := currPt.Z + dz * l.dt
	t := currPt.T + l.dt 

	pt := models.Point {
		X: x, 
		Y: y,
		Z: z,
		T: t,
	}

	return pt 
}

func NewLorenz() *Lorenz {
	return &Lorenz{
		sigma: 10.0,
		rho: 28.0,
		beta: 8.0/3.0,
		init: models.Point {
			X: 0.1,
			Y: 0.0,
			Z: 0.0,
			T: 0.0,
		},
		dt: 0.01,
		numSteps: 10000,
	}
}

func GetLorenzVisualizeContext() vis.VisualizeContext {
	lorenzVis := LorenzVisContext{}
	lorenzVis.vp = vis.VisualizeParams {
			SizeX: 750,
			SizeY: 750,
			IgnoreAxis: "Y",
	}
	
	return lorenzVis
}

func (l LorenzVisContext) ScaleX(x float64) float64 {
	return l.InitX(l.vp.SizeX) + l.ScaleFactor(l.vp.SizeX) * x;
}

func (l LorenzVisContext) ScaleY(y float64) float64 {
	return l.InitY(l.vp.SizeX) + l.ScaleFactor(l.vp.SizeX) * y;
}

func (l LorenzVisContext) ScaleZ(z float64) float64 {
	return l.InitZ(l.vp.SizeX) + l.ScaleFactor(l.vp.SizeX) * z;
}

func (l LorenzVisContext) VisualizeParams() vis.VisualizeParams {
	return l.vp
}

func (l LorenzVisContext) InitX(x float64) float64 {
	return l.vp.SizeX/2
}

func (l LorenzVisContext) InitY(y float64) float64 {
	return l.vp.SizeX/2
}

func (l LorenzVisContext) InitZ(z float64) float64 {
	return 5
}

func (l LorenzVisContext) ScaleFactor(sizeX float64) float64 {
	return sizeX/55
}