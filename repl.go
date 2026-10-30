package main

import (
	"bufio"
	"strings"
	"fmt"
	"os"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		for scanner.Scan() {
			input_str := scanner.Text()
			clean_input := cleanInput(input_str)
			if len(clean_input) > 0 {
				fmt.Printf("Your command was: %s\n", clean_input[0]) 
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
