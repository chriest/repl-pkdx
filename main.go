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
}

func main() {

	cf := con{
		Client: pokeapi.NewClient(time.Minute*5),

	}

	repl(&cf)
}