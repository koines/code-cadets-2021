package solution_test

import (
	"fmt"
	"testing"
	"zad1/solution"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGame(t *testing.T) {
	for idx, tc := range getTestCases() {
		Convey(fmt.Sprintf("Given test case #%v: %+v", idx, tc), t, func() {
			actualOutput, actualError := solution.Game(tc.inputStart, tc.inputEnd)

			if tc.expectingError {
				So(actualError, ShouldNotBeNil)
			} else {
				So(actualError, ShouldBeNil)
				So(actualOutput, ShouldResemble, tc.expectedOutput)
			}

		})
	}
}