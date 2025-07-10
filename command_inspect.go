package main

import "fmt"

//todo

func commandInspect(cfg *config, pokemonName string) error {

	pokemon, exists := cfg.caughtPokemon[pokemonName]

	if !exists {
		fmt.Println("You have not caught this pokemon")
		return nil
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%s: %d\n", stat.StatInfo.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, t := range pokemon.Types {
		fmt.Printf(" -%s\n", t.TypeInfo.Name)
	}
	return nil
}
