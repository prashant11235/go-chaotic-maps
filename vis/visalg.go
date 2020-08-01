package vis

import (
	"math"
	"gochaoticmaps/models"
)

// while sizex, Y & z are separate only square/cube allowed to prevent strethching
func ScalePoints(pts []models.Point, sizeX int, sizeY int, sizeZ int) []models.Point {
	scaledPts := make([]models.Point, len(pts))
	translateX := 0.0
	translateY := 0.0
	translateZ := 0.0

	xpts := PtList(pts, "X")
	ypts := PtList(pts, "Y")
	zpts := PtList(pts, "Z")

	// find min/max of XYZ
	xmin, xmax := MinMax(xpts...)
	ymin, ymax := MinMax(ypts...)
	zmin, zmax := MinMax(zpts...)

	xrange := math.Abs(xmax-xmin)
	yrange := math.Abs(ymax-ymin)
	zrange := math.Abs(zmax-zmin)
	// Perform scaling ie multiply all pts by a const factor such that min is ~0 & max ~sizeX for X, sizeY for Y, size Z
	//_, maxrange := MinMax(math.Abs(xmax-xmin), math.Abs(ymax-ymin), math.Abs(zmax-zmin))
	scaleFactorX := float64(sizeX)/xrange
	scaleFactorY := float64(sizeY)/yrange
	scaleFactorZ := float64(sizeZ)/zrange

	// If min's < 0 - shift all axis by amount (math.Abs(min of x, y, z)) - Now all pts are >0
	if (xmin < 0) {
		translateX = translateX + scaleFactorX*xrange/2
	}
	if (ymin < 0) {
		translateY = translateY + scaleFactorY*yrange/2
	}
	if (zmin < 0) {
		translateZ = translateZ + scaleFactorZ*zrange/2
	}
	
	for idx, pt := range pts {
		scaledPts[idx] = models.Point {
			X: translateX + scaleFactorX * pt.X, 
			Y: translateY + scaleFactorY * pt.Y,
			Z: translateZ + scaleFactorZ * pt.Z,
			T: pt.T,
		}
	}
	
	return scaledPts
}

func PtList(pts []models.Point, dim string) []float64 {
    list := make([]float64, len(pts))
    for _, pt := range pts {
		if dim == "X" {
			list = append(list, pt.X)
		} else if dim == "Y" {
			list = append(list, pt.Y)
		} else {
			list = append(list, pt.Z)
		}
    }
    return list
}

func MinMax(array ...float64) (float64, float64) {
    var max float64 = array[0]
    var min float64 = array[0]
    for _, value := range array[1:] {
        if max < value {
            max = value
        }
        if min > value {
            min = value
        }
    }
    return min, max
}

func GetTranslateFactor(min float64) float64 {
	return min
}