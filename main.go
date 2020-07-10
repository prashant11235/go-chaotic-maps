package main

import (
	"log"
	"github.com/ajstarks/svgo"
	"net/http"
	"fmt"
	"gochaoticmaps/continuous"
	"gochaoticmaps/models"
)

func main() {
	http.Handle("/visualize", http.HandlerFunc(visualize))
	err := http.ListenAndServe(":2003", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func visualize(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	pts := (*continuous.NewLorenz()).GenerateMapPoints()
	path := genPath(pts)

	s := svg.New(w)
	s.Start(1000, 1000)
	//s.Circle(250, 250, 250, "fill:none;stroke:black")
	s.Path(path, "fill:none;stroke:black")
	//s.Path("M150 0 L75 200 L225 200 Z", "fill:none;stroke:black")
	s.End()
}

func genPath(pts []models.Point) (string) {
	path := "M " + fmt.Sprintf("%f", transformX(pts[0].X)) + "," + fmt.Sprintf("%f", transformZ(pts[0].Z)) + " "
	for _, pt := range pts {
		currX := transformX(pt.X)
		currZ := transformZ(pt.Z)
		
		path +=  " L " + fmt.Sprintf("%f", currX) + "," + fmt.Sprintf("%f", currZ)
	}
	return path 
}

func transformX(x float64) float64 {
	return 450 + 20 * x;
}

func transformZ(z float64) float64 {
	return 50 + 20 * z;
}
