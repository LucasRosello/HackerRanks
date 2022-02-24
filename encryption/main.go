package main

import (
	"math"
	"strings"
	"unicode"
)

func encryption(s string) string {
	cleanedString := clean(s)
	upRoot, downRoot := generateDistribution(float64(len(cleanedString)))
	result := encrypt(cleanedString, upRoot, downRoot)
	return result
}

/*
Encript a text
*/
func encrypt(s string, upRoot int32, downRoot int32) string {

	strSlice := make([]string, upRoot)

	index := int32(0)
	for _, letter := range s {

		if index == upRoot {
			index = 0
		}

		strSlice[index] = strSlice[index] + string(letter)

		index++
	}

	return strings.Join(strSlice, " ")
}

/*
Calculates the upper and lower square roots
necessary for the distribution of the letters
*/
func generateDistribution(strLen float64) (int32, int32) {

	root := math.Sqrt(strLen)

	upRoot := math.Ceil(root)    // Raiz superior
	downRoot := math.Floor(root) // Raiz Inferior

	return int32(upRoot), int32(downRoot)
}

/*
Remove spaces from text
*/
func clean(s string) string {
	// Codigo Robado
	rr := make([]rune, 0, len(s))
	for _, r := range s {
		if !unicode.IsSpace(r) {
			rr = append(rr, r)
		}
	}
	return string(rr)
}
