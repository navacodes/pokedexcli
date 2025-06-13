package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()

		input := scanner.Text()

		words := CleanInput(input)

		if len(words) > 0 {
			fmt.Println("Your command was:", words[0])
		}
	}
}
