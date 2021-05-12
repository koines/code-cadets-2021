package taxCalculation_test

import (
	"fmt"
	"testing"
	"zad2/taxCalculation"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCalculateTax(t *testing.T) {
	for idx, tc := range getTestCases() {
		Convey(fmt.Sprintf("Given test case #%v: %+v", idx, tc), t, func() {

			actualOutput, actualErr := taxCalculation.CalculateTax(tc.input)

			if tc.expectingError {
				So(actualErr, ShouldNotBeNil)
			} else {
				So(actualErr, ShouldBeNil)
				So(actualOutput, ShouldResemble, tc.expectedOutput)
			}

		})

	}
}