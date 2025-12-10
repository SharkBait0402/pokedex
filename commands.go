package main

import (
	"os"
	"fmt"
	"github.com/sharkbait0402/pokedex/internal/pokeapi"
	"math/rand"
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

var caughtPokemon = make(map[string]pokeapi.PokemonResponse)

func commandCatch(cfg *config) error {

	pokemon:=cfg.Name
	if pokemon==""{
		return fmt.Errorf("No pokemon was given")
	}

	data, err:=cfg.pokeClient.GetPokemon(pokemon)
	if err!=nil {
		return fmt.Errorf("no pokemon was found with that name")
	}

	fmt.Println("Throwing a Pokeball at " + data.Name + "...")

	chanceCaught:=data.Base_Experience/40

	randomNum:=rand.Intn(chanceCaught + 1)

	chanceStr:=fmt.Sprintf("random: %d, chance: %d", randomNum, chanceCaught)
	fmt.Println(chanceStr)

	if randomNum == chanceCaught {
		fmt.Println(data.Name + " was caught successfully")
		caughtPokemon[data.Name] = data
	}



	return nil

}

func commandInspect(cfg *config) error {
	pokemon:=cfg.Name
	if pokemon == "" {
		return fmt.Errorf("No pokemon was given")
	}

	if data,ok:=caughtPokemon[pokemon]; ok{
		fmt.Println( data)
	} else {
		fmt.Println("pokemon is not caught yet")
	}
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
			"inspect": {
				name: "inspect",
				description: "show pokemon stats if already caught",
				callback: commandInspect,
			},
	}
}
