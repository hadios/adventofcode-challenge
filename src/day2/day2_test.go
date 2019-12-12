package main

import (
	"testing"
)

type testData struct {
	input  []int64
	result []int64
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
			input:  []int64{1, 0, 0, 0, 99},
			result: []int64{2, 0, 0, 0, 99},
		},
		testData{
			input:  []int64{2, 3, 0, 3, 99},
			result: []int64{2, 3, 0, 6, 99},
		},
		testData{
			input:  []int64{2, 4, 4, 5, 99, 0},
			result: []int64{2, 4, 4, 5, 99, 9801},
		},
		testData{
			input:  []int64{1, 1, 1, 4, 99, 5, 6, 0, 99},
			result: []int64{30, 1, 1, 4, 2, 5, 6, 0, 99},
		},
	}

	for _, s := range testDataArr {
		result := executeOpsCommand(s.input)

		if !Equal(result, s.result) {
			t.Errorf("Result was incorrect, got: %d, want: %d.", result, s.result)
		}
	}
}
