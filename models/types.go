package models

type Point struct {
	X float64 
	Y float64
	Z float64
	T float64
}
	
type MapSize struct {
	SizeX float64
	SizeY float64
}

type MapInit struct {
	X float64
	Y float64 
}

type ChaoticMap interface {
    GenerateMapPoints() []Point
}
