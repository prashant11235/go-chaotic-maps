package vis 

import (
	"fmt"
	"gochaoticmaps/models"
	"github.com/ajstarks/svgo"
	"net/http"
)

type PathVisualizer struct {

}

func (pv PathVisualizer) Visualize(mapObj models.ChaoticMap, visctxt VisualizeContext, w http.ResponseWriter) {
	pts := mapObj.GenerateMapPoints()

	path := pv.GenPath(pts, visctxt)

	s := svg.New(w)
	s.Start(int(visctxt.VisualizeParams().SizeX), int(visctxt.VisualizeParams().SizeY))
	s.Path(path, "fill:none;stroke:black;stroke-width:0.2")
	s.End()
}

func (pv PathVisualizer) GenPath(pts []models.Point, visctxt VisualizeContext) (string) {
	path := ""
	if visctxt.VisualizeParams().IgnoreAxis == "X" {
		path += "M " + fmt.Sprintf("%f", visctxt.ScaleY(pts[0].X)) + "," + fmt.Sprintf("%f", visctxt.ScaleZ(pts[0].Z)) + " "
	} else if visctxt.VisualizeParams().IgnoreAxis == "Y" {
		path += "M " + fmt.Sprintf("%f", visctxt.ScaleX(pts[0].X)) + "," + fmt.Sprintf("%f", visctxt.ScaleZ(pts[0].Z)) + " "
	} else if visctxt.VisualizeParams().IgnoreAxis == "Z" {
		path += "M " + fmt.Sprintf("%f", visctxt.ScaleX(pts[0].X)) + "," + fmt.Sprintf("%f", visctxt.ScaleY(pts[0].Z)) + " "
	}
	
	for _, pt := range pts {
		currX := 0.0
		currY := 0.0
		
		if visctxt.VisualizeParams().IgnoreAxis == "X" {
			currX = visctxt.ScaleY(pt.Y)
			currY = visctxt.ScaleZ(pt.Z)
		} else if visctxt.VisualizeParams().IgnoreAxis == "Y" {
			currX = visctxt.ScaleX(pt.X)
			currY = visctxt.ScaleZ(pt.Z)
		} else if visctxt.VisualizeParams().IgnoreAxis == "Z" {
			currX = visctxt.ScaleX(pt.X)
			currY = visctxt.ScaleY(pt.Y)
		}
		
		path +=  " L " + fmt.Sprintf("%f", currX) + "," + fmt.Sprintf("%f", currY)
	}
	return path 
}


func NewPathVisualizer() *PathVisualizer {
	return &PathVisualizer{}
}
