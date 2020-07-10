package models

type Point struct {
	X float64 
	Y float64
	Z float64
}
	
type ChaoticMap interface {
    GenerateMapPoints() []Point
}