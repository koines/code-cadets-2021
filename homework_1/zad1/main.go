package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/koines/code-cadets-2021/homework_1/zad1/fizzbuzz"
)

func main() {
	startPtr := flag.Int("start", 1, "First number")
	endPtr := flag.Int("end", 2, "Last number")
	flag.Parse()

	solution, err := fizzbuzz.Game(*startPtr, *endPtr)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(solution)
}
