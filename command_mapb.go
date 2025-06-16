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
	cfg.previousLocationAreaURL = LocationAreasResponse.Previous

	return nil
}
