package main

import (
	"strconv"
	"strings"
)

func isLucky(n int) bool {
	ticketSlice := strings.Split(strconv.Itoa(n), "")
	firstHalf := ticketSlice[:len(ticketSlice)/2]
	secondHalf := ticketSlice[len(ticketSlice)/2:]

	if sum(firstHalf) == sum(secondHalf) {
		return true
	}
	return false
}

func sum(slice []string) int {
	var count int
	for _, v := range slice {
		value, _ := strconv.Atoi(v)
		count += value
	}

	return count
}
