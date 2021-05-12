package taxCalculation

import (
	"github.com/pkg/errors"
	"math"
)

func CalculateTax(inputValue float64) (float64, error){
	if inputValue < 0 {
		return 0, errors.New("input value is negative")
	}

	var taxLevels = []taxLevel  {
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

	var okLevels bool = checkLevels(taxLevels)

	if !okLevels {
		return 0, errors.New("tax levels are not compatible")
	}

	var result float64 = 0

	for idx, taxLevel := range taxLevels {
		if taxLevel.UpperBound < inputValue {
			if idx != 0 {
				result += (taxLevel.UpperBound-taxLevels[idx-1].UpperBound)*taxLevel.Percentage
			} else {
				result += taxLevel.UpperBound * taxLevel.Percentage
			}
		} else {
			if idx != 0 {
				result += (inputValue - taxLevels[idx-1].UpperBound) * taxLevel.Percentage
			} else {
				result += inputValue * taxLevel.Percentage
			}
			break
		}
	}

	return result, nil
}