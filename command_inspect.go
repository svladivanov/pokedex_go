package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("please provide the name of a caught Pokemon")
	}

	name := args[0]
	pokemon, ok := cfg.caughtPokemon[name]
	if !ok {
		return fmt.Errorf("you have not caught a %s yet", name)
	}

	fmt.Println()
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, ptype := range pokemon.Types {
		fmt.Printf("  - %s\n", ptype.Type.Name)
	}
	fmt.Println()

	return nil
}
