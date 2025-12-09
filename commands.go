package main

import (
	"os"
	"fmt"
	"github.com/sharkbait0402/pokedex/internal/pokeapi"
)

type config struct {
	Next *string
	Previous *string
	pokeClient *pokeapi.Client
	Name string
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _, com := range commands {
		fmt.Printf("%s: %s\n", com.name, com.description)
	}
	return nil
}

func commandMap(cfg *config) error {
	data, _:=cfg.pokeClient.GetLocations(cfg.Next)
	cfg.Next = data.Next
	cfg.Previous = data.Previous
	for i:=0;i<len(data.Results);i++ {
		fmt.Println(data.Results[i].Name)
	}
	return nil
}

func commandMapB(cfg *config) error {


	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	data, _:=cfg.pokeClient.GetLocations(cfg.Previous)
	cfg.Next = data.Next
	cfg.Previous = data.Previous

	for i:=0;i<len(data.Results);i++ {
		fmt.Println(data.Results[i].Name)
	}
	return nil
}

func commandExplore(cfg *config) error {

	location:=cfg.Name
	if location==""{
		return fmt.Errorf("no name was given")
	}

	data, _:=cfg.pokeClient.Explore(location)

	for i:=0;i<len(data.Pokemon_encounters);i++{
		fmt.Println("- ", data.Pokemon_encounters[i].Pokemon.Name)
	}

	cfg.Name = ""
	return nil
}

func commandCatch(cfg *config) error {

	pokemon:=cfg.Name
	if pokemon==""{
		return fmt.Errorf("No pokemon was given")
	}

	data, _:=cfg.pokeClient.GetPokemon(pokemon)

	fmt.Println("Throwing a pokeball at " + data.Name + "...")

	return nil

}

type cliCommand struct {
	name string
	description string
	callback func(*config) error
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
			"explore": {
				name: "explore",
				description: "explore an area",
				callback: commandExplore,
			},
			"catch": {
				name: "catch",
				description: "have a chance to catch a pokemon based on experience level",
				callback: commandCatch,
			},
	}
}
