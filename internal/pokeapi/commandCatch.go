package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"

	"github.com/felixsolom/pokedexcli/internal/pokecache"
)

func CommandCatch(pokemon *PokemonStruct, cache *pokecache.Cache, pokemonName []string) error {
	name := string(pokemonName[0])
	url := baseAPI + "pokemon/" + name

	var body []byte
	entry, exists := cache.Get(url)
	if exists {
		body = entry
	} else {
		res, err := http.Get(url)
		if err != nil {
			return fmt.Errorf("error message: %v", err)
		}

		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)

		if res.StatusCode > 299 {
			return fmt.Errorf("response failed with status code %d, and\nbody %s", res.StatusCode, body)
		}
		if err != nil {
			return fmt.Errorf("error message: %v", err)
		}

		cache.Add(url, body)
	}

	err := json.Unmarshal(body, pokemon)
	if err != nil {
		return fmt.Errorf("error message: %v", err)
	}
	fmt.Printf("Throwing a Pokeball at %v...\n", name)
	const d = 0.02
	prob := 1.0 / (1.0 + d*float64(pokemon.BaseExperience))
	if rand.Float64() < prob {
		fmt.Printf("%v was caught!\n", name)
		AddToPokedex(name, *pokemon)
	} else {
		fmt.Printf("%v escaped!\n", name)
	}
	return nil
}
