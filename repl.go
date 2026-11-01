package main

import (
	"bufio"
	"strings"
	"fmt"
	"os"
	"github.com/pablomotte/pokedex_go/internal/pokeapi"
)

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*config) error
}

type config struct {
	Commands map[string]cliCommand
	BaseUrl string
	NextUrl *string
	PreviousUrl *string
	apiClient pokeapi.Client
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		for scanner.Scan() {
			input_str := scanner.Text()
			clean_input := cleanInput(input_str)
			if len(clean_input) > 0 {
				cmd, ok := cfg.Commands[clean_input[0]]
				if !ok {
					fmt.Printf("Unknown command: %s\n", clean_input[0]) 
				} else {
					if err := cmd.Callback(cfg); err != nil {
						fmt.Printf("Error in %s: %v", cmd.Name, err)
					}
				}
			}
			fmt.Print("Pokedex > ")
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	}
}

func cleanInput(text string) []string {
    return strings.Fields(strings.ToLower(text))
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"map": {
			Name:        "map",
			Description: "Display the next 20 location areas",
			Callback:    commandMap,
		},
		"mapp": {
			Name:        "mapp",
			Description: "Display the previous 20 location areas",
			Callback:    commandMapP,
		},
		"help": {
			Name:        "help",
			Description: "Display the Pokedex help text",
			Callback:    commandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
	}
}
