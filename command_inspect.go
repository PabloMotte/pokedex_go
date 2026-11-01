package main

import (
	"fmt"
)

func commandInspect(cfg *config) error {
	if cfg.parms == nil || len(cfg.parms) < 1 {
		fmt.Printf("%spokemon name or id# required%s\n", ColorRed, ColorNone)
		return nil
	}
	pokemon, ok := cfg.pokedex[cfg.parms[0]]
	if !ok {
		fmt.Printf("%s%s is not in your Pokedex, you must catch it first!%s\n", ColorRed, poke_name(cfg.parms[0]), ColorNone)
		return nil
	}

	fmt.Printf("%s%s was caught!%s\n", ColorBlue, poke_name(pokemon.Name), ColorNone)
	fmt.Printf("%sName: %s%s\n", ColorBlue, ColorNone, poke_name(pokemon.Name))
	fmt.Printf("%sHeight: %s%d\n", ColorBlue, ColorNone, pokemon.Height)
	fmt.Printf("%sWeight: %s%d\n", ColorBlue, ColorNone, pokemon.Weight)
	fmt.Printf("%sStats:%s\n", ColorBlue, ColorNone)
	for _, stat := range pokemon.Stats {
		fmt.Printf("%s  -%s: %s%d\n", ColorBlue, stat.Stat.Name, ColorNone, stat.BaseStat)
	}
	fmt.Printf("%sTypes:%s\n", ColorBlue, ColorNone)
	for _, t := range pokemon.Types {
		fmt.Printf("%s  - %s%s\n", ColorBlue, ColorNone, t.Type.Name)
	}

	return nil
}
