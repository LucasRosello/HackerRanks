package main

import "testing"

type Test struct {
	in  [][]int
	out int
}

var testsCases = []Test{
	{
		// 1
		[][]int{{}},
		0,
	},
	{
		// 2
		[][]int{{1}},
		1,
	},
	{
		// 3
		[][]int{{1, 1, 1, 1}},
		4,
	},
	{
		// 4
		[][]int{{0}, {1}},
		0,
	},
	{
		// 5
		[][]int{{1, 1, 1, 0}, {0, 5, 0, 1}, {2, 1, 3, 10}},
		9,
	},
}

func TestMatrixElementsSum(t *testing.T) {

	for i, test := range testsCases {
		result := matrixElementsSum(test.in)
		if result != test.out {
			t.Errorf("Error in the test N°%v, Expected %v, but got %v. For the input %v", i+1, test.out, result, test.in)
		}
	}

}
