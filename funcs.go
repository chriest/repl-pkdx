package main

import (
	//"github.com/chriest/repl-pkdx/internal/pokeapi"
	"strings"
	"fmt"
	"os"
	"errors"
	"math/rand"
)



func formatter(input string) []string{
	re := strings.TrimSpace(input)
	re = strings.ToLower(re)
	red := strings.Fields(re)
	return red
}
  
func mapCommand(cf *con, args ...string) error{
	/*if cf.Succesive == nil {
		return errors.New("last page")
	}*/
	resp, err := cf.Client.ListAreas(cf.Succesive)
	if err!= nil {
		fmt.Println("error")
	}
	fmt.Println("Areas: ")

	for _, item := range resp.Results {
		fmt.Print(" - ")
		fmt.Println(item.Name)
	}
	cf.Succesive = resp.Next
	cf.Previous = resp.Previous
	return nil
}

func mapBCommand(cf *con, args ...string) error{
	if cf.Previous == nil {
		return errors.New("first page")
	}
	resp, err := cf.Client.ListAreas(cf.Previous)
	if err!= nil {
		fmt.Println("error")
	}
	fmt.Println("Areas: ")

	for _, item := range resp.Results {
		fmt.Print(" - ")
		fmt.Println(item.Name)
	}
	cf.Succesive = resp.Next
	cf.Previous = resp.Previous
	return nil
}

func exitCommand(cf *con, args ...string) error{
	//fmt.Println("exit function")
	os.Exit(0)
	return nil
}

func explore(cf *con, args ...string) error{
	if len(args)!=1{
		return errors.New("err")
	}
	locationAreaName:=args[0]

	locationArea, err := cf.Client.GetAreas(locationAreaName)
	if err!= nil {
		fmt.Println("error")
	}
	fmt.Printf("In %v:", locationArea.Name)
	fmt.Println("git ")

	for _, item := range locationArea.PokemonEncounters {
		fmt.Print(" - ")
		fmt.Println(item.Pokemon.Name)
	}
	return nil

}

func inspect(cf *con, args ...string) error{
	if len(args)!=1{
		return errors.New("err")
	}
	pokemonName:=args[0]

	pokemon, err := cf.catchers[pokemonName]

	if !err {
		return errors.New("pokemon not caught")
	}
	
	fmt.Printf("Name %v\n", pokemon.Name)
	fmt.Printf("Height %v\n", pokemon.Height)
	fmt.Printf("Weight %v\n", pokemon.Weight)
	/*for _, stt := range pokemon.Stat {
	fmt.Printf(" %v: %v", stt.Stat.Name, stt.Stat.BaseStat)
	}*/


	return nil

}

func pokedex(cf *con, args ...string) error{
	fmt.Println("Printing Pokedex: ")

	for _, pokemon := range cf.catchers{
		fmt.Println(pokemon.Name)
	}


	return nil

}

func catch(cf *con, args ...string) error{
	if len(args)!=1{
		return errors.New("err")
	}
	pokemonName:=args[0]

	pokemon, err := cf.Client.GetPokemons(pokemonName)
	if err!= nil {
		fmt.Println("error")
	}
	threshold := 50
	randGen := rand.Intn(pokemon.BaseExperience)
	if randGen < threshold {
		cf.catchers[pokemonName] = pokemon
		fmt.Printf("Pokemon caught: %v\n", pokemonName)
	} else {
		return fmt.Errorf("pokemon not caught")
	}

	return nil

}

func helpCommand(cf *con, args ...string) error{
	fmt.Printf("Welcome to the Pokedex!\n")
	fmt.Printf("Usage:\n")
	for _, item := range getCommands() {
		fmt.Printf(" * %v : %v\n", item.name, item.description)
	}
	return nil
}