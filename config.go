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
