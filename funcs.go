package main

import (
	"strings"
	"fmt"
	"os"
)



func formatter(input string) string{
	re := strings.TrimSpace(input)
	re = strings.ToLower(re)
	return re
}
  
func mapCommand() error{
	return nil
}

func mapBCommand() error{
	return nil
}

func exitCommand() error{
	fmt.Println("exit function")
	os.Exit(0)
	return nil
}


func helpCommand() error{
	fmt.Printf("Welcome to the Pokedex!\n")
	fmt.Printf("Usage:\n")
	for _, item := range getCommands() {
		fmt.Printf(" * %v : %v\n", item.name, item.description)
	}
	return nil
}