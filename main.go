package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
	"time"
	"github.com/sharkbait0402/pokedex/internal/pokeapi"
) 

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	client:=*pokeapi.NewClient(5*time.Second, 5*time.Second)

	cfg:= config{
		pokeClient: &client,
	}
	
	for ;; {
		fmt.Print("Pokedex > ")

		ok:=scanner.Scan()
		if !ok {
			break
		}

		input:=strings.ToLower(scanner.Text())

		if input == "" {
			fmt.Println("Please enter a command")
			continue
		}

		comm:=commandWord(input)
		if comm == "explore" {
			cfg.Name=restInput(input)
			if cmd, ok:=commands[comm]; ok {
				fmt.Println("exploring now...")
				if err:=cmd.callback(&cfg); err!=nil{
					fmt.Println("error: ", err)
				}
			}
			
		}else if comm == "catch" {

			cfg.Name=restInput(input)
			if cmd, ok:=commands[comm]; ok{
				if err:=cmd.callback(&cfg); err!=nil {
					fmt.Println("error: ", err)
				}
			}

		}else if cmd, ok := commands[comm]; ok {
			if err:=cmd.callback(&cfg); err!=nil {
				fmt.Println("error:", err)
			}
		}
		

	}

	return
}
