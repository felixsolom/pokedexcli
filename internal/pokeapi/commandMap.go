package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func CommandMap(config *Config) error {
	url := config.Next
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area"
	}
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
	err = json.Unmarshal(body, config)
	if err != nil {
		return fmt.Errorf("error message: %v", err)
	}
	for _, area := range config.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func CommandMapb(config *Config) error {
	url := config.Previous
	if url == "" {
		fmt.Println("you're on the first page")
		return nil
	}
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
	err = json.Unmarshal(body, config)
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
