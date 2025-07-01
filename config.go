package main

import "github.com/navacodes/pokedexcli/internal/pokecache"

type config struct {
	nextLocationAreaURL     *string
	previousLocationAreaURL *string
	cache                   *pokecache.Cache
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}
type LocationAreasResponse struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Ecounters - Pokemon encounters structs , these retrive the pokemont located at the LocationArea
type Encounters struct {
	PokemonEncounters []Pokemon `json:"pokemon_encounters"`
}

type Pokemon struct {
	Pokemon PokemonInfo `json:"pokemon"`
}
type PokemonInfo struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
