package main

import (
	"os"
	"fmt"
	"github.com/sharkbait0402/pokedex/internal/pokeapi"
)

var cfg struct {
	Next *string
	Previous *string
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	for _, com := range commands {
		fmt.Printf("%s: %s\n", com.name, com.description)
	}
	return nil
}

func commandMap() error {
	data, _:=pokeapi.GetLocations(cfg.Next)
	cfg.Next = data.Next
	cfg.Previous = data.Previous
	for i:=0;i<len(data.Results);i++ {
		fmt.Println(data.Results[i].Name)
	}
	return nil
}

func commandMapB() error {

	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	data, _:=pokeapi.GetLocations(cfg.Previous)
	cfg.Next = data.Next
	cfg.Previous = data.Previous

	for i:=0;i<len(data.Results);i++ {
		fmt.Println(data.Results[i].Name)
	}
	return nil
}

type cliCommand struct {
	name string
	description string
	callback func() error
}

var commands map[string]cliCommand

func init() {

	commands = map[string]cliCommand {
			"help": {
				name: "help",
				description: "Display help message",
				callback: commandHelp,
			},
			"exit": {
				name: "exit",
				description: "Exit the Pokedex",
				callback: commandExit,
			},
			"map": {
				name: "map",
				description: "show 20 locations starting with first page",
				callback: commandMap,
			},
			"mapb": {
				name: "mapb",
				description: "show previous page of locations",
				callback: commandMapB,
			},
	}
}
