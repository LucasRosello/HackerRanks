package main

import "testing"

type Test struct {
	n       int
	queens  [][]int
	queries [][]int
	out     []bool
}

var testsCases = []Test{
	{
		1,
		[][]int{{0, 1}},
		[][]int{{0, 1}},
		[]bool{false},
	},
}

func TestSquaresUnderQueenAttack(t *testing.T) {

	for i, test := range testsCases {
		result := squaresUnderQueenAttack(test.n, test.queens, test.queries)
		if !ThisSlicesAreTheSame(result, test.out) {
			t.Errorf("Error in the test N°%v, Expected %v, but got %v.", i+1, test.out, result)
		}
	}

}

func ThisSlicesAreTheSame(a, b []bool) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestDiagonal(t *testing.T) {
	queen := []int{1, 1}
	query := []int{2, 0}

	result := CheckDiagonalPoint(queen, query)

	if result != true {
		t.Errorf("Not equal")
	}
}
