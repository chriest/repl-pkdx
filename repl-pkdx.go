package main

import (
	"fmt"
	"bufio"
	"os"
)

type cliMethods struct {
	name 		string
	description	string
	callback 	func() error
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
			"mapb": {
				name:			"mapb",
				description:	"Shows the 20 locations before",
				callback:		mapBCommand,
			},
		}
}

func repl() {
		inBuf := bufio.NewScanner(os.Stdin);
		input := ""
		
		
		fmt.Print("pokedex > ")
		for inBuf.Scan(){
			input = formatter(inBuf.Text())
			if comm, errNone := getCommands()[input]; errNone {
				e := comm.callback()
				if e != nil {
					fmt.Println(e)
				}
			} else {
				fmt.Println("incorrect or empty command")
			}
			fmt.Print("pokedex > ")
		}
}

func main(){
	repl()
}