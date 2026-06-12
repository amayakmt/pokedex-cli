package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/amayakmt/pokedex-cli/internal/pokeapi"
)

func main() {
	commands := getCommands()
	cfg := &config{
		Client: pokeapi.NewClient(),
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())

		if len(words) == 0 {
			continue
		}

		command, ok := commands[words[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}

	}
}

// Helper commands ---------------------------------------

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	Next     *string
	Previous *string
	Client   pokeapi.Client
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: `Displays the names of 20 location areas in the Pokemon world`,
			callback:    commandMap,
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
