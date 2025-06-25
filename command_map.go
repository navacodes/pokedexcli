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

	//check cache first
	var response LocationAreasResponse
	if cachedData, found := cfg.cache.Get(url); found {
		fmt.Println("Using cached data!")
		// Parse the cached response instead of making HTTP request
		if err := json.Unmarshal(cachedData, &response); err != nil {
			fmt.Println("Error unmarshaling cached data:", err)
			return err
		}

		for _, location := range response.Results {
			fmt.Println(location.Name)
		}

		cfg.nextLocationAreaURL = response.Next
		cfg.previousLocationAreaURL = response.Previous
		return nil
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

	cfg.cache.Add(url, body)

	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Println("Error Unmarshal:", err)
	}

	for _, location := range response.Results {
		fmt.Println(location.Name)
	}

	cfg.nextLocationAreaURL = response.Next
	cfg.previousLocationAreaURL = response.Previous
	return nil
}
