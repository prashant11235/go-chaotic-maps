package vis 

import (
	"gochaoticmaps/models"
	"github.com/ajstarks/svgo"
	"net/http"
)

type CircleVisualizer struct {

}

func (cv CircleVisualizer) Visualize(mapObj models.ChaoticMap, visctxt VisualizeContext, w http.ResponseWriter) {
	pts := mapObj.GenerateMapPoints()
	scaledPts := ScalePoints(pts, visctxt.VisualizeParams().Size.SizeX, 
		visctxt.VisualizeParams().Size.SizeY, visctxt.VisualizeParams().Size.SizeZ)

	s := svg.New(w)
	s.Start(int(visctxt.VisualizeParams().Size.SizeX), int(visctxt.VisualizeParams().Size.SizeY))
	cv.GenCircles(scaledPts, visctxt, s)
	s.End()
}

func (cv CircleVisualizer) GenCircles(pts []models.Point, visctxt VisualizeContext, s *svg.SVG) {
	for _, pt := range pts {
		currX := pt.X
		currY := pt.Y
		s.Circle(int(currX), int(currY), 1, "fill:none;stroke:black")
	}
}


func NewCircleVisualizer() *CircleVisualizer {
	return &CircleVisualizer{}
}
