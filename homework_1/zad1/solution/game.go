package solution

import (
	"github.com/pkg/errors"
	"strconv"
)

func Game(first int, last int ) ([]string, error){
	if first > last {
		return nil, errors.New("start number is greater then end number")
	}

	var result []string

	for i := first; i <= last; i++ {
		if i % 15 == 0 {
			result = append(result, "FizzBuzz")
		} else if i % 3 == 0 {
			result = append(result, "Fizz")
		} else if i % 5 == 0 {
			result = append(result, "Buzz")
		} else {
			result = append(result, strconv.Itoa(i) )
		}
	}

	return result, nil
}