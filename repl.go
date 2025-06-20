package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/felixsolom/pokedexcli/internal/pokeapi"
	"github.com/felixsolom/pokedexcli/internal/pokecache"
)

type cliCommamd struct {
	name        string
	description string
	callback    func(args []string) error
}

func getCommands(config *pokeapi.Config, cache *pokecache.Cache) map[string]cliCommamd {
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
			callback: func(args []string) error {
				return pokeapi.CommandMap(config, cache, args)
			},
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 location areas",
			callback: func(args []string) error {
				return pokeapi.CommandMapb(config, cache, args)
			},
		},
		"explore": {
			name:        "explore",
			description: "Displays all pokemons found in specified area",
			callback: func(args []string) error {
				location := &pokeapi.LocationAreaNameID{}
				return pokeapi.CommandExplore(location, cache, args)
			},
		},
		"catch": {
			name:        "catch",
			description: "Attempting to catch a specific pokemon",
			callback: func(args []string) error {
				name := &pokeapi.PokemonStruct{}
				return pokeapi.CommandCatch(name, cache, args)
			},
		},
	}
}

func startRepl(config *pokeapi.Config, cache *pokecache.Cache) {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands(config, cache)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		commandWord := input[0]
		args := []string{}
		if len(input) > 1 {
			args = input[1:]
		}
		command, exists := commands[commandWord]
		if exists {
			err := command.callback(args)
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

func cleanInput(text string) []string {
	lower_str := strings.ToLower(text)
	sliced_strs := strings.Fields(lower_str)
	return sliced_strs
}

func commandExit(args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(args []string) error {
	output := ""
	for _, entry := range getCommands(nil, nil) {
		output += fmt.Sprintf("%s: %s\n", entry.name, entry.description)
	}
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	fmt.Println(output)
	return nil
}
