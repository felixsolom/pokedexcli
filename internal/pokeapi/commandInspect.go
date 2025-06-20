package pokeapi

import "fmt"

func CommandInspect(args []string) error {
	name := args[0]
	pokemon, exists := GetFromPokedex(name)
	if exists {
		fmt.Printf("Name: %v\n", pokemon.Name)
		fmt.Printf("Height: %v\n", pokemon.Height)
		fmt.Printf("Weight: %v\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("-%v: ", stat.Stat.Name)
			fmt.Printf("%v\n", stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, t := range pokemon.Types {
			fmt.Printf("-%v\n", t.Type.Name)
		}
		return nil
	}
	return fmt.Errorf("this pokemon is still roaming free")
}
