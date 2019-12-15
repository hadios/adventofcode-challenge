package main

import "testing"

type testData struct {
	input  int64
	result int64
}

func TestCalculateFuel(t *testing.T) {

	/*
	** Part 1
	 */
	testDataArr := []testData{
		testData{
			input:  12,
			result: 2,
		},
		testData{
			input:  14,
			result: 2,
		},
		testData{
			input:  1969,
			result: 654,
		},
		testData{
			input:  100756,
			result: 33583,
		},
	}

	for _, s := range testDataArr {
		total := calculateFuel(s.input)

		if total != s.result {
			t.Errorf("Result was incorrect, got: %d, want: %d.", total, s.result)
		}
	}

	/*
	** Part 2
	 */
	testDataArr2 := []testData{
		testData{
			input:  14,
			result: 2,
		},
		testData{
			input:  1969,
			result: 966,
		},
	}

	for _, s := range testDataArr2 {
		total := calculateFuelRepeat(s.input)

		if total != s.result {
			t.Errorf("Result was incorrect, got: %d, want: %d.", total, s.result)
		}
	}
}
