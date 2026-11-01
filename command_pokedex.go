package main

import (
	"fmt"
)

func commandPokedex(cfg *config) error {
	if cfg.pokedex == nil || len(cfg.pokedex) < 1 {
		fmt.Printf("%sno pokemon have been captured%s\n", ColorBlue, ColorNone)
		return nil
	}

	fmt.Printf("%sYour Pokedex:%s\n", ColorBlue, ColorNone)
	for k, _ := range cfg.pokedex {
		fmt.Printf("%s  - %s%s\n", ColorBlue, ColorNone, poke_name(k))
	}

	return nil
}
