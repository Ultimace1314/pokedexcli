package main

import "fmt"

func commandInspect(cft *config, pokemon ...string) error {
	if len(pokemon) != 1 {
		return fmt.Errorf("you must provide the name of the pokemon to catch")
	}
	pokemonName := pokemon[0]
	pokemonData, ok := cft.pokedex[pokemonName]
	if !ok {
		return fmt.Errorf("you have not caught %s", pokemonName)
	}
	// print pokemon data
	fmt.Printf("Name: %s\n", pokemonData.Name)
	// print pokemon id
	fmt.Printf("ID: %d\n", pokemonData.ID)
	fmt.Printf("Height: %d\n", pokemonData.Height)
	fmt.Printf("Weight: %d\n", pokemonData.Weight)
	// print pokemon stats
	fmt.Printf("Stats:\n")
	for _, stat := range pokemonData.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	// print pokemon types
	fmt.Printf("Types:\n")
	for _, ptype := range pokemonData.Types {
		fmt.Printf("  -%s\n", ptype.Type.Name)
	}
	// print pokemon abilities
	fmt.Printf("Abilities:\n")
	for _, ability := range pokemonData.Abilities {
		fmt.Printf("  -%s\n", ability.Ability.Name)
	}

	return nil
}
