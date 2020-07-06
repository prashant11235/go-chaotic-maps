package continuous

type Point struct {
	X float64 
	Y float64
	Z float64
}

var maxSteps = 5000

// Constants
var sigma = float64(10)
var rho = float64(28)
var beta = 8.0 / 3.0

// Initial values
var x = 0.1
var y = 0.1
var z = 0.1

// Time
var dt = 0.01


func GenerateMapPoints() ([]Point) {
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
	dx := sigma * (y - x)
	dy := x * (rho - z) - y
	dz := x * y - beta * z

	x += dx * dt
	y += dy * dt 
	z += dz * dt

	pt := Point {
		X: x, 
		Y: y,
		Z: z,
	}

	return pt 
}
