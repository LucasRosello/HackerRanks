package main

func matrixElementsSum(matrix [][]int) int {

	var sum int
	var hauntedNumbers = make(map[int]bool)

	for topIndex, slice := range matrix {
		for i := 0; i < len(slice); i++ {
			if !hauntedNumbers[i] {
				if matrix[topIndex][i] == 0 {
					hauntedNumbers[i] = true
				}
				sum += matrix[topIndex][i]
			}
		}
	}

	return sum
}
