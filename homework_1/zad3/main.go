package main

import (
	"log"
	"os"
	"strings"
	"zad3/pokemon"
)

func main()  {
	if len(os.Args) != 2 {
		log.Fatalln("wrong input: needed to be name or number of Pokemon")
	}

	input := os.Args[1]
	input = strings.ToLower(input)

	pokemon.FindPokemonLocations(input)
}
