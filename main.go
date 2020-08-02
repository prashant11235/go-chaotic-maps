package main

import (
	"log"
	"strings"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"gochaoticmaps/continuous"
	"gochaoticmaps/discrete"
	"gochaoticmaps/models"
	"gochaoticmaps/vis"
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
	draw(w, r, mapName)
}

func draw(w http.ResponseWriter, req *http.Request, mapName string) {
	w.Header().Set("Content-Type", "image/svg+xml")
	mapObj, visctxt, err := instantiate(mapName)
	if err != nil {
		log.Fatal(err)
	} else {

		// Generate visualization object
		pathVis := *vis.NewCircleVisualizer()
		pathVis.Visualize(mapObj, visctxt, w)
	}
}

func instantiate(mapName string) (models.ChaoticMap, vis.VisualizeContext, error) {
	switch strings.ToLower(mapName) {
	case "lorenz":
		return *continuous.NewLorenz(), vis.GetDefaultVisualizeContext(), nil
	case "duffing":
		return *discrete.NewDuffing(), vis.GetDefaultVisualizeContext(), nil
	case "duffingcont":
		return *continuous.NewDuffing(), vis.GetDefaultVisualizeContext(), nil
	case "thomas":
		return *continuous.NewThomas(), vis.GetDefaultVisualizeContext(), nil
	case "henon":
		return *discrete.NewHenon(), vis.GetDefaultVisualizeContext(), nil
	case "kaplanyorke":
		return *discrete.NewKaplanYorke(), vis.GetDefaultVisualizeContext(), nil
	case "peterdejong":
		return *discrete.NewPeterDeJong(), vis.GetDefaultVisualizeContext(), nil
	default:
		log.Fatal("Can't process !!")
		return nil, nil, errors.New("Can't process")
	}
}