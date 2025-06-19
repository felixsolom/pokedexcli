package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/felixsolom/pokedexcli/internal/pokeapi"
	"github.com/felixsolom/pokedexcli/internal/pokecache"
)

func getCommands() map[string]cliCommamd {
	return map[string]cliCommamd{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays next 20 location areas",
			callback:    pokeapi.CommandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 location areas",
			callback:    pokeapi.CommandMapb,
		},
	}
}

func startRepl(config *pokeapi.Config, cache *pokecache.Cache) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		commandWord := input[0]
		command, exists := getCommands()[commandWord]
		if exists {
			err := command.callback(config, cache)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

type cliCommamd struct {
	name        string
	description string
	callback    func(*pokeapi.Config, *pokecache.Cache) error
}

func cleanInput(text string) []string {
	lower_str := strings.ToLower(text)
	sliced_strs := strings.Fields(lower_str)
	return sliced_strs
}

func commandExit(_ *pokeapi.Config, _ *pokecache.Cache) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(_ *pokeapi.Config, _ *pokecache.Cache) error {
	output := ""
	for _, entry := range getCommands() {
		output += fmt.Sprintf("%s: %s\n", entry.name, entry.description)
	}
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	fmt.Println(output)
	return nil
}
