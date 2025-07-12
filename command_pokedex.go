package main

import "fmt"

func commandPokedex(cfg *config, areaName string) error {
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("You have not caught any pokemon")
		return nil
	}
	fmt.Println("Your Pokedex:")

	//loop through caughtPokemon and print them out
	//map[string]PokemonCaught
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf("- %s\n", pokemon.Name)
	}

	return nil
}
