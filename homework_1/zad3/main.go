package main

import (
	"log"
	"os"
	"strings"

	"github.com/koines/code-cadets-2021/homework_1/zad3/pokemon"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("wrong input: needed to be name or number of Pokemon")
	}

	input := os.Args[1]
	input = strings.ToLower(input)

	err := pokemon.FindPokemonLocations(input)
	if err != nil {
		log.Fatal(err)
	}
}
