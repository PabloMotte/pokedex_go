package main

import (
	"bufio"
	"strings"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}


func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display the Pokedex help text",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		for scanner.Scan() {
			input_str := scanner.Text()
			clean_input := cleanInput(input_str)
			if len(clean_input) > 0 {
				cmd, ok := getCommands()[clean_input[0]]
				if !ok {
					fmt.Printf("Unknown command: %s\n", clean_input[0]) 
				} else {
					if err := cmd.callback(); err != nil {
						fmt.Printf("Error in %s: %v", cmd.name, err)
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
