package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"
)

func commandCatch(cfg *config, pokemonName string) error {
	if pokemonName == "" {
		fmt.Println("You cant catch blank, provide a Pokemon Name you want to catch")
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	//TODO:
	//1. Build API URL for pokemon endpoint
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName
	//2. check cache / make http reques
	var response PokemonCatch

	if cachedData, found := cfg.cache.Get(url); found {
		fmt.Println("Using cached Data")
		if err := json.Unmarshal(cachedData, &response); err != nil {
			return err
		}
	} else {

		res, err := http.Get(url)
		if err != nil {
			fmt.Println("error:", err)
			return err
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		cfg.cache.Add(url, body)

		if err := json.Unmarshal(body, &response); err != nil {
			return err
		}

	}
	// catch formula

	PokemonExperience := response.BaseExperience

	catchChance := max(5, 60-(PokemonExperience/4))
	randomNum := rand.IntN(100)

	if randomNum < catchChance {
		fmt.Printf("%s was caught!\n", pokemonName)
		// need add to the pokemon map storage
		cfg.caughtPokemon[pokemonName] = response

	} else {
		fmt.Printf("%s escaped! \n", pokemonName)
	}

	return nil
}
