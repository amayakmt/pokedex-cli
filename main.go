package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/amayakmt/pokedex-cli/internal/pokeapi"
)

func main() {
	commands := getCommands()
	cfg := &config{
		Client:  pokeapi.NewClient(5 * time.Minute),
		Pokedex: make(map[string]pokeapi.Pokemon),
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

		cfg.Args = words[1:]
		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	Next     *string
	Previous *string
	Client   pokeapi.Client
	Args     []string
	Pokedex  map[string]pokeapi.Pokemon
}
