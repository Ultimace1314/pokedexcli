package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func getPokemonLocations(url string) pokeLocations {

	res, err := http.Get(url)

	if err != nil {
		return pokeLocations{}
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return pokeLocations{}
	}
	if err != nil {
		return pokeLocations{}
	}

	locations := pokeLocations{}
	er := json.Unmarshal(body, &locations)
	if er != nil {
		return pokeLocations{}
	}
	return locations

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
