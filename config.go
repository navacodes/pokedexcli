package main

type config struct {
	nextLocationAreaURL     *string
	previousLocationAreaURL *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}
