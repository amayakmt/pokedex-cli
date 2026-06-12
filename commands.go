package main

import (
	"fmt"
	"math/rand"
	"os"
)

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Displays all the pokemons in the area. Syntax: explore <area-name>",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch the pokemon. Syntax: catch <pokemon-name>",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Shows details of a Pokemon if it has been caught. Syntax: inspect <pokemon-name>",
			callback:    commandInspect,
		},
	}
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Printf("Welcome to the Pokedex!\n")

	fmt.Printf("Usage:\n\n")
	for _, command := range getCommands() {
		fmt.Printf("%v: %v\n", command.name, command.description)
	}

	return nil
}

func commandMap(cfg *config) error {
	resp, err := cfg.Client.GetLocationAreas(cfg.Next)
	if err != nil {
		return err
	}

	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}

	cfg.Next = resp.Next
	cfg.Previous = resp.Previous

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	resp, err := cfg.Client.GetLocationAreas(cfg.Previous)
	if err != nil {
		return err
	}

	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}

	cfg.Next = resp.Next
	cfg.Previous = resp.Previous

	return nil
}

func commandExplore(cfg *config) error {
	if len(cfg.Args) == 0 {
		return fmt.Errorf("Please provide a location name")
	}
	name := cfg.Args[0]

	resp, err := cfg.Client.GetLocationArea(name)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", name)
	fmt.Println("Found Pokemon:")
	for _, pok := range resp.PokemonEncounter {
		fmt.Printf(" - %s\n", pok.Pokemon.Name)
	}

	return nil
}

func commandCatch(cfg *config) error {
	if len(cfg.Args) == 0 {
		return fmt.Errorf("Please provide a pokemon name")
	}
	name := cfg.Args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	pokemon, err := cfg.Client.GetPokemon(name)
	if err != nil {
		return err
	}

	// higher base experience = harder to catch
	if rand.Intn(pokemon.BaseExperience) > 40 {
		fmt.Printf("%s escaped!\n", name)
		return nil
	}

	fmt.Printf("%s was caught!\n", name)
	cfg.Pokedex[name] = pokemon
	return nil
}

func commandInspect(cfg *config) error {
	if len(cfg.Args) == 0 {
		return fmt.Errorf("Please provide a Pokemon name")
	}
	name := cfg.Args[0]

	pokemon, ok := cfg.Pokedex[name]
	if !ok {
		fmt.Println("You have not caught that Pokemon")
		return nil
	}

	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)

	fmt.Println("Stats:")
	for _, s := range pokemon.Stats {
		fmt.Printf("  - %v: %v\n", s.Stat.Name, s.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  %v\n", t.Type.Name)
	}

	return nil
}
