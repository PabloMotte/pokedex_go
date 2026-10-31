package main

import (
	"bufio"
	"strings"
	"fmt"
	"os"
	"github.com/pablomotte/pokedex_go/internal/commands"
)

func startRepl(cfg *commands.Config) {
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
