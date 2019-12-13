package main

import (
	"testing"
)

type testData struct {
	input  []int64
	result []int64
}

type testData2 struct {
	input  []int64
	input2 []int64
	result int64
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
			input:  []int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
			result: []int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
		},
		testData{
			input:  []int64{1102, 34915192, 34915192, 7, 4, 7, 99, 0},
			result: []int64{1219070632396864},
		},
		testData{
			input:  []int64{104, 1125899906842624, 99},
			result: []int64{1125899906842624},
		},
	}

	for _, s := range testDataArr {
		_, result := executeOpsCommand(s.input, []int64{})

		if !Equal(result, s.result) {
			t.Errorf("Result was incorrect, got: %d, want: %d.", result, s.result)
		}
	}

	/*
	** Test cases from previous days to ensure other opcode still works
	 */
	testDataArr2 := []testData2{
		// input == 8 ? 1 : 0
		testData2{
			input:  []int64{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
			input2: []int64{1},
			result: 0,
		},
		testData2{
			input:  []int64{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
			input2: []int64{8},
			result: 1,
		},

		// input < 8 ? 1 : 0
		testData2{
			input:  []int64{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
			input2: []int64{1},
			result: 1,
		},
		testData2{
			input:  []int64{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
			input2: []int64{8},
			result: 0,
		},

		// input == 8 ? 1 : 0
		testData2{
			input:  []int64{3, 3, 1108, -1, 8, 3, 4, 3, 99},
			input2: []int64{8},
			result: 1,
		},
		testData2{
			input:  []int64{3, 3, 1108, -1, 8, 3, 4, 3, 99},
			input2: []int64{9},
			result: 0,
		},

		// input < 8 ? 1 : 0
		testData2{
			input:  []int64{3, 3, 1107, -1, 8, 3, 4, 3, 99},
			input2: []int64{7},
			result: 1,
		},
		testData2{
			input:  []int64{3, 3, 1107, -1, 8, 3, 4, 3, 99},
			input2: []int64{9},
			result: 0,
		},

		// input == 0 ? 0 : 1
		testData2{
			input:  []int64{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
			input2: []int64{0},
			result: 0,
		},
		testData2{
			input:  []int64{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
			input2: []int64{8},
			result: 1,
		},

		// num < 8, 999
		// num == 8, 1000
		// num > 8, 1001
		testData2{
			input: []int64{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
				1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
				999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			input2: []int64{6},
			result: 999,
		},
		testData2{
			input: []int64{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
				1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
				999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			input2: []int64{8},
			result: 1000,
		},
		testData2{
			input: []int64{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
				1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
				999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			input2: []int64{9},
			result: 1001,
		},
	}

	for _, s := range testDataArr2 {
		_, output := executeOpsCommand(s.input, s.input2)

		if output[0] != s.result {
			t.Errorf("Using test %d \nResult was incorrect, got: %d, want: %d.", s, output, s.result)
		}
	}
}
