package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandExplore(cfg *config, areaName string) error {
	if areaName == "" {
		fmt.Println("Input an area to explore")
		return nil
	}

	url := "https://pokeapi.co/api/v2/location-area/" + areaName
	var response Encounters

	// Get data either from cache or HTTP request
	if cachedData, found := cfg.cache.Get(url); found {
		fmt.Println("Using cached Data")
		if err := json.Unmarshal(cachedData, &response); err != nil {
			fmt.Println("Error unmarshaling cached data:", err)
			return err
		}
	} else {
		// Cache miss - make HTTP request
		res, err := http.Get(url)
		if err != nil {
			fmt.Println("error:", err)
			return err
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println("error:", err)
			return err
		}

		cfg.cache.Add(url, body)

		if err := json.Unmarshal(body, &response); err != nil {
			fmt.Println("Error Unmarshal:", err)
			return err
		}
	}

	for _, pokemonInfo := range response.PokemonEncounters {
		fmt.Println(pokemonInfo.Pokemon.Name)
	}

	return nil
}
