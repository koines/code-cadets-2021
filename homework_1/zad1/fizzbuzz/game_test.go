package fizzbuzz_test

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/koines/code-cadets-2021/homework_1/zad1/fizzbuzz"
)

func TestGame(t *testing.T) {
	for idx, tc := range getTestCases() {
		Convey(fmt.Sprintf("Given test case #%v: %+v", idx, tc), t, func() {
			actualOutput, actualError := fizzbuzz.Game(tc.inputStart, tc.inputEnd)

			if tc.expectingError {
				So(actualError, ShouldNotBeNil)
			} else {
				So(actualError, ShouldBeNil)
				So(actualOutput, ShouldResemble, tc.expectedOutput)
			}

		})
	}
}
