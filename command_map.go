package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMap(cfg *config) error {

	var url string

	if cfg.nextLocationAreaURL == nil {
		url = "https://pokeapi.co/api/v2/location-area"

	} else {
		url = *cfg.nextLocationAreaURL
	}
	res, err := http.Get(url)

	if err != nil {
		fmt.Println("error:", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error:", err)
	}

	var LocationAreasResponse LocationAreasResponse
	if err := json.Unmarshal(body, &LocationAreasResponse); err != nil {
		fmt.Println("Error Unmarshal:", err)
	}

	for _, location := range LocationAreasResponse.Results {
		fmt.Println(location.Name)
	}

	cfg.nextLocationAreaURL = LocationAreasResponse.Next
	cfg.nextLocationAreaURL = LocationAreasResponse.Previous
	return nil
}
