package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func apicall(url string, config *config) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	Body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	locations := pokeLocations{}
	err = json.Unmarshal(Body, &locations)
	if err != nil {
		return err
	}
	config.nextLocationsUrl = locations.Next
	config.previousLocationsUrl = locations.Previous
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	return nil
}

type pokeLocations struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
