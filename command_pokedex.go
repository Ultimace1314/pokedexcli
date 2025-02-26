package main

func commandPokedex(cfg *config, args ...string) error {
	// Print the names of all the pokemon in the pokedex
	if len(cfg.pokedex) == 0 {
		println("Your pokedex is empty!")
		return nil
	}
	println("Your pokedex:")
	for name := range cfg.pokedex {
		println(" - ", name)
	}
	return nil
}
