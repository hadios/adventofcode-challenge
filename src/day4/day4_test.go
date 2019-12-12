package main

import (
	"testing"
)

type testData struct {
	input  int64
	result bool
}

func Equal(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func Test(t *testing.T) {
	/*
	** Part 1
	 */
	testDataArr := []testData{
		testData{
			input:  111111,
			result: true,
		},
		testData{
			input:  223450,
			result: false,
		},
		testData{
			input:  123789,
			result: false,
		},
	}

	for _, s := range testDataArr {
		result := isIncreasingDigits(s.input) && hasDoubleDigits(s.input)

		if result != s.result {
			t.Errorf("Result was incorrect, got: %t, want: %t.", result, s.result)
		}
	}

	/*
	** Part 2
	 */
	testDataArr2 := []testData{
		testData{
			input:  112233,
			result: true,
		},
		testData{
			input:  123444,
			result: false,
		},
		testData{
			input:  111122,
			result: true,
		},
		testData{
			input:  123455,
			result: true,
		},
		testData{
			input:  222333,
			result: false,
		},
		testData{
			input:  188999,
			result: true,
		},
	}

	for _, s := range testDataArr2 {
		result := isIncreasingDigits(s.input) && hasDoubleDigits(s.input) && hasNoLargerDoubleDigits(s.input)

		if result != s.result {
			t.Errorf("Result was incorrect, got: %t, want: %t.", result, s.result)
		}
	}
}
