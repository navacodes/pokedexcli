package main

import "github.com/navacodes/pokedexcli/internal/pokecache"

type config struct {
	nextLocationAreaURL     *string
	previousLocationAreaURL *string
	cache                   *pokecache.Cache
	caughtPokemon           map[string]PokemonCatch
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
type PokemonCatch struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []Stat `json:"stats"`
	Types          []Type `json:"types"`
}
type Stat struct {
	BaseStat int      `json:"base_stat"`
	StatInfo StatInfo `json:"stat"`
}
type StatInfo struct {
	Name string `json:"name"`
}
type Type struct {
	TypeInfo TypeInfo `json:"type"`
}
type TypeInfo struct {
	Name string `json:"name"`
}
