package main

import (
	"time"

	"github.com/felixsolom/pokedexcli/internal/pokeapi"
	"github.com/felixsolom/pokedexcli/internal/pokecache"
)

func main() {
	cache := pokecache.NewCache(5 * time.Second)
	config := &pokeapi.Config{}
	startRepl(config, cache)
	cache.Shutdown()
}
