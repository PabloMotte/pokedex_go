package commands

import (
	"fmt"
)

func commandHelp(cfg *Config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println()
	// fmt.Println("help: Displays a help message\nexit: Exit the Pokedex")
	for _, cmd := range cfg.Commands {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}
	return nil
}
