package main

import "fmt"

func main() {}

func squaresUnderQueenAttack(n int, queens [][]int, queries [][]int) []bool {

	for _, v := range queens {
		fmt.Printf("La reina esta posicionada en [ x = %d y = %d ]\n ", v[0], v[1])
	}

	return []bool{false}
}

func CheckDiagonalPoint(queen, query []int) bool {
	return queen[0]-query[0] == queen[1]-query[1]
}
