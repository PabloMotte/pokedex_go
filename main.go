package main

import "github.com/pablomotte/pokedex_go/internal/commands"


func main() {
	s_location_area := "location-area"
	cfg := &commands.Config{
		Commands: commands.GetCommands(),
		BaseUrl: "https://pokeapi.co/api/v2/",
		NextUrl: &s_location_area,
		PreviousUrl: nil,
	}
	startRepl(cfg)
	return
}
