package solution_test

type testCase struct {
	inputStart int
	inputEnd int

	expectedOutput []string
	expectingError bool
}

func getTestCases() []testCase {
	return []testCase {
		{
			inputStart: 1,
			inputEnd: 10,

			expectedOutput: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz"},
			expectingError: false,
		},
		{
			inputStart: 10,
			inputEnd: 15,

			expectedOutput: []string{"Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
			expectingError: false,
		},
		{
			inputStart: 2,
			inputEnd: 1,

			expectingError: true,
		},
	}
}