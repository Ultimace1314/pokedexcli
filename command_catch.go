package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cft *config, pokemon ...string) error {

	if len(pokemon) != 1 {
		return fmt.Errorf("you must provide the name of the pokemon to catch")
	}
	pokemonName := pokemon[0]
	pokemonData, err := cft.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	chance := rand.Intn(pokemonData.BaseExperience)
	if chance > 50 {
		fmt.Printf("%s escaped!\n", pokemonName)
	} else {
		fmt.Printf("%s was caught!\n", pokemonName)
		cft.pokedex[pokemonName] = pokemonData
	}

	return nil
}
