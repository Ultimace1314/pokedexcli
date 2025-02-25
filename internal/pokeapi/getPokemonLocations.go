package pokeapi

import (
	"fmt"
	"io"
	"net/http"
)

func getPokemonLocations(url string) (pokeLocations, error) {
	res, err := http.Get(url)
	if err != nil {
		return pokeLocations{}, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return pokeLocations{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	return pokeLocations{}, nil
}

type pokeLocations struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
