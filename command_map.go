package main

import (
	github.com/Ultimace1314/internal/pokeapi
)

func commandMap(config *config) ([]string, error) {
	if config.nextLocationsUrl != nil {
		return getPokemonLocations(*config.nextLocationsUrl), nil
	} else {
		return getPokemonLocations("https://pokeapi.co/api/v2/location-area/"), nil
	}

}
