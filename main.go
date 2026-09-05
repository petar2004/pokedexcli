package main

import (
	"pokedex/internal/pokecache"
	"time"
)

func main() {

	cfg := &config{
		next:    "https://www.pokeapi.co/api/v2/location-area/",
		cache:   pokecache.NewCache(5 * time.Second),
		pokedex: make(map[string]pokemonDetails),
	}

	cfg.commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays a list of maps",
			callback:    commandMap,
		},
		"explore": {
			name:        "explore",
			description: "Displays a list of all the Pokémon located",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a Pokeman",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a Pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all caught pokemons",
			callback:    commandPokedex,
		},
	}

	startRepl(cfg)
}
