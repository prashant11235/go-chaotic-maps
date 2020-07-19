package main

import (
	"log"
	"strings"
	"errors"
	"github.com/gorilla/mux"
)


func instantiate(mapName string) (models.ChaoticMap, error) {
	switch strings.ToLower(mapName) {
	case "lorenz":
		return *continuous.NewLorenz(), nil
	case "duffing":
		return *discrete.NewDuffing(), nil
	case "duffingcont":
		return *continuous.NewDuffing(), nil
	default:
		log.Fatal("Can't process !!")
		return nil, errors.New("Can't process")
	}
}