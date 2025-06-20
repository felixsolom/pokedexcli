package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/felixsolom/pokedexcli/internal/pokecache"
)

func CommandExplore(names *LocationAreaNameID, cache *pokecache.Cache, LocationName []string) error {
	name := string(LocationName[0])
	url := baseAPI + "location-area/" + name

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

	err := json.Unmarshal(body, names)
	if err != nil {
		return fmt.Errorf("error message: %v", err)
	}
	fmt.Printf("exploring %s...\n", LocationName)
	for _, encounter := range names.PokemonEncounters {
		fmt.Printf("- %v\n", encounter.Pokemon.Name)
	}
	return nil
}
