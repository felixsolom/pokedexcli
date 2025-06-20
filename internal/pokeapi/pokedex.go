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

func GetFromPokedex(name string) (PokemonStruct, bool) {
	pokedexMu.RLock()
	defer pokedexMu.RUnlock()
	pokemon, exists := pokedex[name]
	return pokemon, exists
}
