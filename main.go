package main

import (
	"log"
	"github.com/gorilla/mux"
	"github.com/ajstarks/svgo"
	"net/http"
	"gochaoticmaps/continuous"
	"gochaoticmaps/draw"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/visualize/{map}", VisualizeHandler)
	err := http.ListenAndServe(":2003", r)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func VisualizeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mapName := vars["map"]
	w.WriteHeader(http.StatusOK)
	visualize(w, r, mapName)
}

func visualize(w http.ResponseWriter, req *http.Request, mapName string) {
	w.Header().Set("Content-Type", "image/svg+xml")
	pts := (*continuous.NewLorenz()).GenerateMapPoints()

	draw := *draw.NewDraw()
	path := draw.GenPath(pts)

	s := svg.New(w)
	s.Start(draw.SizeX, draw.SizeY)
	s.Path(path, "fill:none;stroke:black")
	s.End()
}
