package main

import (
	"github.com/chriest/repl-pkdx/internal/pokeapi"
	"time"
	//"errors"
)

type con struct {
	Client pokeapi.Client
	Previous *string
	Succesive *string
	catchers map[string]pokeapi.Pokemons
}

func main() {

	cf := con{
		Client: pokeapi.NewClient(time.Minute*5),
		catchers: make(map[string]pokeapi.Pokemons),
	}

	repl(&cf)
}