package discrete 

import (
	"math"
)

// Constants
var alpha = float64(1)
var beta = float64(-1)
var delta = float64(0.2)
var gamma = float64(0.3)
var omega = float64(1)
var maxSteps = 2500

// Initial values
var x = 0.0
var y = 0.0
var t = 0.0

// Time variation
var dt = 0.04

type Point struct {
	X float64 
	Y float64
}

func GeneratePoints() ([]Point) {
	i := 0 
	ptArr := make([]Point, maxSteps)
	for i < maxSteps {
		pt := calcNextPoint()
		ptArr[i] = pt
		i++
	} 

	return ptArr 

}

func calcNextPoint() (Point) {
	dx := float64(y)
	dy := -delta*float64(y) - beta*float64(x) - alpha*float64(x*x*x) + gamma*(float64(math.Cos(omega*t)))
	
	x += dx * dt
	y += dy * dt
	
	t += dt
	
	pt := Point {
		X: x, 
		Y: y,
	}

	return pt 
}
