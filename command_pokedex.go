package main

import (
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error {

	if len(cfg.caughtPokemon) == 0 {

		fmt.Println("You don't have any pokemon in your pokedex")

		return nil
	}
	fmt.Println("Pokemon in Pokedex:")

	for _, pokemon := range cfg.caughtPokemon {

		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil

}
