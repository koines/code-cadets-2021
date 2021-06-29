package taxCalculation

import (
	"github.com/pkg/errors"
)

func CalculateTax(inputValue float64, taxLevels []TaxLevel) (float64, error) {
	if inputValue < 0 {
		return 0, errors.New("input value is negative")
	}

	var okLevels bool = checkLevels(taxLevels)
	if !okLevels {
		return 0, errors.New("tax levels are not compatible")
	}

	var result float64 = 0

	for idx, taxLevel := range taxLevels {
		if taxLevel.UpperBound < inputValue {
			if idx != 0 {
				result += (taxLevel.UpperBound - taxLevels[idx-1].UpperBound) * taxLevel.Percentage
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
