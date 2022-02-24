package main

import "testing"

func TestIsLucky(t *testing.T) {
	type Test struct {
		in  int
		out bool
	}

	var testCases = []Test{
		{
			12,
			false,
		},
		{
			11,
			true,
		},
	}

	for _, test := range testCases {
		result := isLucky(test.in)
		if result != test.out {
			t.Errorf("Expected %v, got %v for the input %v", test.out, result, test.in)
		}
	}
}
