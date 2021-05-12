package taxCalculation_test

type testCase struct {
	input float64

	expectedOutput float64
	expectingError bool
}

func getTestCases() []testCase {
	return []testCase {
		{
			input: 7000,

			expectedOutput: 800,
			expectingError: false,
		},
		{
			input: 250,

			expectedOutput: 0,
			expectingError: false,
		},
		{
			input: -2500,

			expectingError: true,
		},
	}
}
