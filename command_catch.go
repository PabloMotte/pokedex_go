package main

import (
	"fmt"
	"math/rand"
	"time"
	"github.com/pablomotte/pokedex_go/internal/pokeapi"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func commandCatch(cfg *config) error {
	if cfg.BaseUrl == nil {
		fmt.Printf("no base URL has been set\n")
		return nil
	}
	s := *cfg.BaseUrl + "pokemon"
	err := print_catch_pokemon(s, cfg)
	if err != nil {
		return err
	}
	return nil
}

func print_catch_pokemon(url string, cfg *config) error {
	// use int64(99) for a "fixed" random number (first is always 73 in my testing), 
	// or current nanoseconds for complete pseudo-randomness
	// seed := int64(99) // time.Now().UnixNano() 
	seed := time.Now().UnixNano() 
	if cfg.parms == nil || len(cfg.parms) < 1 {
		fmt.Printf("%spokemon name or id# required%s\n", ColorRed, ColorNone)
		return nil
	}
	if p, ok := cfg.pokedex[cfg.parms[0]]; ok {
		fmt.Printf("%s%s is already in your Pokedex, you can only have one!%s\n", ColorRed, poke_name(p.Name), ColorNone)
		return nil
	}
	fmt.Printf("%sThrowing a Pokeball at %s...%s\n", ColorBlue, cfg.parms[0], ColorNone)
	jsonBlob, err := cfg.apiClient.GetResponseById(&url, cfg.parms[0])
	if err != nil {
		return err
	}
	if len(jsonBlob) == 0 {
		return nil
	}
	pokemon, err := pokeapi.GetPokemon(jsonBlob)
	if err != nil {
		return err
	}
	if pokemon.BaseExperience > 255 {
		fmt.Printf("%s%s knocks the Pokeball away, derisively %s(%d)%s\n", ColorBlue, poke_name(pokemon.Name), ColorYellow, pokemon.BaseExperience, ColorNone)
		return nil
	}
	r := rand.New(rand.NewSource(seed))
	// check := r.Float64() // dec [0, 1)
	check := r.Intn(256) // int [0, 255)
	if pokemon.BaseExperience > check {
		fmt.Printf("%s%s escaped! %s(%d vs. %d)%s\n", ColorBlue, poke_name(pokemon.Name), ColorYellow, check, pokemon.BaseExperience, ColorNone)
		return nil
	}
	cfg.pokedex[pokemon.Name] = &pokemon
	fmt.Printf("%s%s was caught! %s(%d vs. %d)%s\n", ColorBlue, poke_name(pokemon.Name), ColorYellow, check, pokemon.BaseExperience, ColorNone)
	fmt.Printf("%sYou may now inspect %s with the inspect command.%s\n", ColorYellow, poke_name(pokemon.Name), ColorNone)
	return nil
}

func poke_name(name string) string {
	c := cases.Caser(cases.Title(language.Und, cases.NoLower))
	return c.String(name)
}
