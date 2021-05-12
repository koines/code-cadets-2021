package main

import (
	"flag"
	"fmt"
	"log"
	"zad1/solution"
)

func main() {
	startPtr := flag.Int("start",1,"First number")
	endPtr := flag.Int("end", 2, "Last number")
	flag.Parse()

	solution, err := solution.Game(*startPtr, *endPtr)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(solution)
}
