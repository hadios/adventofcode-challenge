package main

import (
	"testing"
)

type testData struct {
	input  []string
	result int
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
			input:  []string{"R8,U5,L5,D3", "U7,R6,D4,L4"},
			result: 6,
		},
		testData{
			input:  []string{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83"},
			result: 159,
		},
		testData{
			input:  []string{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"},
			result: 135,
		},
	}

	for _, s := range testDataArr {
		result := drawWires(s.input)

		if result != s.result {
			t.Errorf("Result was incorrect, got: %d, want: %d.", result, s.result)
		}
	}

	/*
	** PART 2
	 */
	testDataArr2 := []testData{
		testData{
			input:  []string{"R8,U5,L5,D3", "U7,R6,D4,L4"},
			result: 30,
		},
		testData{
			input:  []string{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83"},
			result: 610,
		},
		testData{
			input:  []string{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"},
			result: 410,
		},
	}

	for _, s := range testDataArr2 {
		result := drawSignalWires(s.input)

		if result != s.result {
			t.Errorf("Result was incorrect, got: %d, want: %d.", result, s.result)
		}
	}
}
