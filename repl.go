package main

import (
	"bufio"
	"strings"
	"fmt"
	"os"
	"github.com/pablomotte/pokedex_go/internal/pokeapi"
)

const (
	ColorRed = "\033[0;31m"
	ColorGreen = "\033[0;32m"
	ColorYellow = "\033[0;33m"
	ColorBlue = "\033[0;36m"
	ColorNone = "\033[0m"
)

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*config) error
}

type config struct {
	Commands map[string]cliCommand
	BaseUrl *string
	NextUrl *string
	PreviousUrl *string
	parms []string
	apiClient pokeapi.Client
	pokedex map[string]*pokeapi.Pokemon
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		printPrompt()
		for scanner.Scan() {
			input_str := scanner.Text()
			cfg.parms = make([]string, 0)
			clean_input := cleanInput(input_str)
			if len(clean_input) > 0 {
				if len(clean_input) > 1 {
					cfg.parms = clean_input[1:]
				}
				cmd, ok := cfg.Commands[clean_input[0]]
				if !ok {
					fmt.Printf("Unknown command: %s\n", clean_input[0]) 
				} else {
					if err := cmd.Callback(cfg); err != nil {
						fmt.Printf("Error in %s: %v", cmd.Name, err)
					}
				}
			}
			printPrompt()
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	}
}

func printPrompt() {
	fmt.Printf("%sPokedex > %s", ColorGreen, ColorNone)
}

func cleanInput(text string) []string {
    return strings.Fields(strings.ToLower(text))
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"catch": {
			Name:        "catch",
			Description: "Catch the specified pokemon and add to your pokedex",
			Callback:    commandCatch,
		},
		"explore": {
			Name:        "explore",
			Description: "Display the pokemon in the specified location area",
			Callback:    commandExplore,
		},
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
