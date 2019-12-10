package main

import (
	"testing"
)

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

func TestCalculateFuel(t *testing.T) {
	var initial = []string{"R8,U5,L5,D3", "U7,R6,D4,L4"}
	var ans = 6
	var total = drawWires(initial)

	if total != ans {
		t.Errorf("Result was incorrect, got: %d, want: %d.", total, ans)
	}

	initial = []string{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83"}
	ans = 159
	total = drawWires(initial)

	if total != ans {
		t.Errorf("Result was incorrect, got: %d, want: %d.", total, ans)
	}

	initial = []string{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"}
	ans = 135
	total = drawWires(initial)

	if total != ans {
		t.Errorf("Result was incorrect, got: %d, want: %d.", total, ans)
	}

	/*
	** PART 2
	 */
	initial = []string{"R8,U5,L5,D3", "U7,R6,D4,L4"}
	ans = 30
	total = drawSignalWires(initial)

	if total != ans {
		t.Errorf("Result was incorrect, got: %d, want: %d.", total, ans)
	}

	initial = []string{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83"}
	ans = 610
	total = drawSignalWires(initial)

	if total != ans {
		t.Errorf("Result was incorrect, got: %d, want: %d.", total, ans)
	}

	initial = []string{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"}
	ans = 410
	total = drawSignalWires(initial)

	if total != ans {
		t.Errorf("Result was incorrect, got: %d, want: %d.", total, ans)
	}
}
