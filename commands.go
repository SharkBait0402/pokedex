package main

import (
	"os"
	"fmt"
)

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
	}
}
