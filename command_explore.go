package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// TODO:
func commandExplore(cfg *config, areaName string) error {
	if areaName == "" {
		fmt.Println("Input an area to explore")
		return nil
	}
	url := "https://pokeapi.co/api/v2/location-area/" + areaName
	var response Encounters

	if cachedData, found := cfg.cache.Get(url); found {
		fmt.Println("Using cached Data")

		//parse the cached response

		if err := json.Unmarshal(cachedData, &response); err != nil {
			fmt.Println("Error unmarshaling cached data:", err)
			return err
		}
		for _, pokemonInfo := range response.PokemonEncounters {
			fmt.Println(pokemonInfo.Pokemon.Name)
		}

		return nil
	}
	//make http request if no cache

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("error:", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println("error:", err)
	}

	cfg.cache.Add(url, body)

	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Println("Error Unmarshal:", err)
	}

	for _, pokemonInfo := range response.PokemonEncounters {
		fmt.Println(pokemonInfo.Pokemon.Name)
	}

	return nil
}
