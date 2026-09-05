package main

import "pokedex/internal/pokecache"

type config struct {
	commands map[string]cliCommand
	next     string
	previous string
	cache    *pokecache.Cache
	pokedex  map[string]pokemonDetails
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type locationAreaResponse struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []locationArea `json:"results"`
}

type locationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type locationAreaDetails struct {
	PokemonEncounters []pokemonEncounters `json:"pokemon_encounters"`
}

type pokemonEncounters struct {
	Pokemon namedResource `json:"pokemon"`
}

type pokemonDetails struct {
	Name           string             `json:"name"`
	BaseExperience int                `json:"base_experience"`
	Height         int                `json:"height"`
	IsDefault      bool               `json:"is_default"`
	Order          int                `json:"order"`
	Weight         int                `json:"weight"`
	Abilities      []pokemonAbilities `json:"abilities"`
	Stats          []pokemonStats     `json:"stats"`
	Types          []types            `json:"types"`
}

type pokemonAbilities struct {
	IsHidden bool          `json:"is_hidden"`
	Slot     int           `json:"slot"`
	Ability  namedResource `json:"ability"`
}

type pokemonStats struct {
	BaseStat int           `json:"base_stat"`
	Effort   int           `json:"effort"`
	Stat     namedResource `json:"stat"`
}

type types struct {
	Slot int           `json:"slot"`
	Type namedResource `json:"type"`
}

type namedResource struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
