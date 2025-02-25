package main

func commandMap(config *config) ([]string, error) {
	if config.nextLocationsUrl != nil {
		return getPokemonLocations(*config.nextLocationsUrl), nil
	} else {
		return getPokemonLocations("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"), nil
	}

}
