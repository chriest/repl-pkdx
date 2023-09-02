package main

import (
	"fmt"
	"bufio"
	"os"
)

type cliMethods struct {
	name 		string
	description	string
	callback 	func(*con, ...string) error
}

func getCommands() map[string]cliMethods {
		return map[string]cliMethods{
			// look for methods in funcs.go
			"help": {
				name:			"help",
				description:	"Help message",
				callback:		helpCommand,
			},
			"exit": {
				name: 			"exit",
				description:	"Exit the Pokedex",
				callback:		exitCommand,
			},
			"map": {
				name:			"map",
				description:	"Show the next 20 locations",
				callback:		mapCommand,
			},
			"explore": {
				name: "explore",
				description: "Pokemons in your area",
				callback: explore,
			},
			"mapb": {
				name:			"mapb",
				description:	"Shows the 20 locations before",
				callback:		mapBCommand,
			},
		}
}

func repl(configure *con) {
		inBuf := bufio.NewScanner(os.Stdin);
		input := ""
		
		
		fmt.Print("pokedex > ")
		args:=[]string{}
		for inBuf.Scan(){
			clean:=formatter(inBuf.Text())
			if len(clean)>1{
				args=clean[1:]
			}
			input = clean[0]
			if comm, errNone := getCommands()[input]; errNone {
				e  := comm.callback(configure, args...)
				if e != nil {
					fmt.Println(e)
				}
			} else {
				fmt.Println("incorrect or empty command")
			}
			fmt.Print("pokedex > ")
		}
}