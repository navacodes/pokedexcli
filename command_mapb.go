package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMapb(cfg *config) error {
	var url string

	if cfg.previousLocationAreaURL == nil {
		fmt.Println("you're on the first page")
		return nil

	}
	url = *cfg.previousLocationAreaURL

	//check for cache first

	var response LocationAreasResponse
	if cachedData, found := cfg.cache.Get(url); found {
		fmt.Println("Using cached data!")

		//Parse the cached response
		if err := json.Unmarshal(cachedData, &response); err != nil {
			return err
		}
		for _, location := range response.Results {
			fmt.Println(location.Name)
		}

		cfg.nextLocationAreaURL = response.Next
		cfg.previousLocationAreaURL = response.Previous
		return nil
	}

	//no cache -- make HTTP request
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
