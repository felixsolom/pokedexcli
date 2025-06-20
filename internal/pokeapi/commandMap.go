package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/felixsolom/pokedexcli/internal/pokecache"
)

type Config struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func CommandMap(config *Config, cache *pokecache.Cache, _ []string) error {
	url := config.Next
	if url == "" {
		url = baseAPI + "location-area"
	}

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

	err := json.Unmarshal(body, config)
	if err != nil {
		return fmt.Errorf("error message: %v", err)
	}
	for _, area := range config.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func CommandMapb(config *Config, cache *pokecache.Cache, _ []string) error {
	url := config.Previous
	if url == "" {
		fmt.Println("you're on the first page")
		return nil
	}
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

		body, err := io.ReadAll(res.Body)

		if res.StatusCode > 299 {
			return fmt.Errorf("response failed with status code %d, and\nbody %s", res.StatusCode, body)
		}
		if err != nil {
			return fmt.Errorf("error message: %v", err)
		}

		cache.Add(url, body)

	}

	err := json.Unmarshal(body, config)
	if err != nil {
		return fmt.Errorf("error message: %v", err)
	}

	if config.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	for _, area := range config.Results {
		fmt.Println(area.Name)
	}
	return nil
}
