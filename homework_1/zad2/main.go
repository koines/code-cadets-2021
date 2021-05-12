package main

import (
	"fmt"
	"log"
	"zad2/taxCalculation"
)

func main() {
	fmt.Println("Enter your income:")

	var money float64
	var currency string

	fmt.Scanf("%f %s", &money, &currency)

	fmt.Println("Calculating...")
	solution, err := taxCalculation.CalculateTax(money)
	if err != nil {
		log.Fatal(err)
	}

	if currency != "" {
		fmt.Println(solution, currency)
	} else {
		fmt.Println(solution)
	}
}