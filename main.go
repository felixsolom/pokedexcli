package main

import (
	"time"

	"github.com/felixsolom/pokedexcli/internal/pokeapi"
	"github.com/felixsolom/pokedexcli/internal/pokecache"
)

func main() {
	config := &pokeapi.Config{}
	startRepl(config)
	pokecache.NewCache(5 * time.Second)
}
