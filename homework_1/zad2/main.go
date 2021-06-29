package main

import (
	"fmt"
	"log"
	"math"

	"github.com/koines/code-cadets-2021/homework_1/zad2/taxCalculation"
)

func main() {
	fmt.Println("Enter your income:")

	var money float64
	var currency string

	_, err := fmt.Scanf("%f %s", &money, &currency)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Calculating...")
	var taxLevels = []taxCalculation.TaxLevel{
		{
			1000,
			0,
		},
		{
			5000,
			0.1,
		},
		{
			10000,
			0.2,
		},
		{
			math.Inf(1),
			0.3,
		},
	}
	solution, err := taxCalculation.CalculateTax(money, taxLevels)
	if err != nil {
		log.Fatal(err)
	}

	if currency != "" {
		fmt.Println(solution, currency)
	} else {
		fmt.Println(solution)
	}
}
