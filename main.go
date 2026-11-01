package main

import (
	"time"
	"github.com/pablomotte/pokedex_go/internal/pokeapi"
)

func main() {
	s_base_url := "https://pokeapi.co/api/v2/"
	s_location_area := "location-area"
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		Commands: GetCommands(),
		BaseUrl: &s_base_url,
		NextUrl: &s_location_area,
		PreviousUrl: nil,
		apiClient: pokeClient,
		pokedex: map[string]*pokeapi.Pokemon{},
	}
	startRepl(cfg)
}
