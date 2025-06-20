package pokeapi

import "fmt"

func CommandPokedex(_ []string) error {
	fmt.Println("Your Pokedex:")
	for _, name := range pokedex {
		fmt.Printf(" - %v\n", name.Name)
	}
	return nil
}
