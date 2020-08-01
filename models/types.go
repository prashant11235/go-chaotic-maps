package models

type Point struct {
	X float64 
	Y float64
	Z float64
	T float64
}
	
type MapSize struct {
	SizeX int
	SizeY int
	SizeZ int
}

type ChaoticMap interface {
    GenerateMapPoints() []Point
}
