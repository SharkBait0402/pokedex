package main

import (
	"os"
	"fmt"
)

func commandExit() error {
	fmt.Println("Closing the pokedex... Goodbye!")
	os.Exit(0)
}

type cliCommand struct {
	name string
	description string
	callback func() error
}

map[string]cliCommand{
    "exit": {
        name:        "exit",
        description: "Exit the Pokedex",
        callback:    commandExit,
    },
}
