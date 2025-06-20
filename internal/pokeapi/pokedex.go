package pokeapi

import "sync"

var (
	pokedex   = make(map[string]PokemonStruct)
	pokedexMu sync.RWMutex
)

func AddToPokedex(name string, pokemon PokemonStruct) {
	pokedexMu.Lock()
	defer pokedexMu.Unlock()
	pokedex[name] = pokemon
}
