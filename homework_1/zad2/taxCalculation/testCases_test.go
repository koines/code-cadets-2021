package taxCalculation_test

import (
	"github.com/koines/code-cadets-2021/homework_1/zad2/taxCalculation"
	"math"
)

type testCase struct {
	input          float64
	inputTaxLevels []taxCalculation.TaxLevel

	expectedOutput float64
	expectingError bool
}

func getTestCases() []testCase {
	return []testCase{
		{
			input: 7000,
			inputTaxLevels: []taxCalculation.TaxLevel{
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
			},

			expectedOutput: 800,
			expectingError: false,
		},
		{
			input: 250,
			inputTaxLevels: []taxCalculation.TaxLevel{
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
			},

			expectedOutput: 0,
			expectingError: false,
		},
		{
			input: -2500,
			inputTaxLevels: []taxCalculation.TaxLevel{
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
			},

			expectingError: true,
		},
		{
			input: 2500,
			inputTaxLevels: []taxCalculation.TaxLevel{
				{
					1000,
					0,
				},
				{
					5000,
					0.1,
				},
				{
					4000,
					0.2,
				},
				{
					math.Inf(1),
					0.3,
				},
			},

			expectingError: true,
		},
		{
			input: 1500,
			inputTaxLevels: []taxCalculation.TaxLevel{
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
					0.15,
				},
			},

			expectingError: true,
		},
	}
}
