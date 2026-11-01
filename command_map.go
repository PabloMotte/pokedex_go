package main

import (
	"fmt"
	"strings"
	"github.com/pablomotte/pokedex_go/internal/pokeapi"
)

func commandMap(cfg *config) error {
	if cfg.NextUrl == nil {
		fmt.Printf("you're on the last page\n")
		return nil
	}
	s := *cfg.NextUrl
	if !strings.HasPrefix(s, "http") {
		s = *cfg.BaseUrl + s
	}
	err := print_locations(s, cfg)
	if err != nil {
		return err
	}
	return nil
}

func print_locations(url string, cfg *config) error {
	jsonBlob, err := cfg.apiClient.GetResponseById(&url, "")
	if err != nil {
		return err
	}
	if len(jsonBlob) == 0 {
		return nil
	}
	locations, err := pokeapi.GetLocationAreas(jsonBlob)
	if err != nil {
		return err
	}
	cfg.NextUrl = locations.NextUrl
	cfg.PreviousUrl = locations.PreviousUrl
	for _, area := range locations.Results {
		fmt.Println(area.Name)
	}
	return nil
}
