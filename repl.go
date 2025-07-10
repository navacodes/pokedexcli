package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		input := reader.Text()
		words := CleanInput(input)

		if len(words) == 0 {
			continue
		}
		commandName := words[0] // grab the command
		cmd, exists := getCommands()[commandName]

		var areaName string

		if len(words) > 1 {
			areaName = words[1]
		} else {
			areaName = ""
		}

		if exists {
			err := cmd.callback(cfg, areaName)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func CleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)

	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display previous 20 location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Display pokemons in area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "attempt to catch pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "pokemon stats and information",
			callback:    commandInspect,
		},
	}

}
