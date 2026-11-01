package main

import (
	"fmt"
	"github.com/pablomotte/pokedex_go/internal/pokeapi"
)

func commandExplore(cfg *config) error {
	if cfg.BaseUrl == nil {
		fmt.Printf("no base URL has been set\n")
		return nil
	}
	s := *cfg.BaseUrl + "location-area"
	err := print_location_pokemon(s, cfg)
	if err != nil {
		return err
	}
	return nil
}

func print_location_pokemon(url string, cfg *config) error {
	if cfg.parms == nil || len(cfg.parms) < 1 {
		fmt.Printf("%slocation name or id# required%s\n", ColorRed, ColorNone)
		return nil
	}
	fmt.Printf("%sExploring %s%s\n", ColorBlue, cfg.parms[0], ColorNone)
	jsonBlob, err := cfg.apiClient.GetResponseById(&url, cfg.parms[0])
	if err != nil {
		return err
	}
	if len(jsonBlob) == 0 {
		return nil
	}
	location_area, err := pokeapi.GetLocationArea(jsonBlob)
	if err != nil {
		return err
	}
	fmt.Printf("%sFound Pokemon: %d%s\n", ColorBlue, len(location_area.PokemonEncounters), ColorNone)
	for _, area := range location_area.PokemonEncounters {
		fmt.Println(area.Pokemon.Name)
	}
	return nil
}
