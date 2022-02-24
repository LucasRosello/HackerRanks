package main

import (
	"reflect"
	"testing"
)

func TestCommonCharacterCounters(t *testing.T) {
	type Test struct {
		s1  string
		s2  string
		out int
	}

	var testsCases = []Test{
		{
			// 1
			"abc",
			"def",
			0,
		},
		// {
		// 	// 2
		// 	"abc",
		// 	"aef",
		// 	1,
		// },
	}

	for i, test := range testsCases {
		result := commonCharacterCount(test.s1, test.s2)
		if result != test.out {
			t.Errorf("Error in the test N°%v, Expected %v, but got %v. For the input %v and %v", i+1, test.out, result, test.s1, test.s2)
		}
	}
}

func TestCountLetters(t *testing.T) {

	type Test struct {
		in  string
		out map[string]int
	}

	var testsCases = []Test{
		{
			"a",
			map[string]int{"a": 1},
		},
		{
			"abbbc",
			map[string]int{"a": 1, "b": 3, "c": 1},
		},
	}

	for i, test := range testsCases {
		result := countLetters(test.in)
		if !reflect.DeepEqual(result, test.out) {
			t.Errorf("Error in the test N°%v, Expected %v, but got %v. For the input %v", i+1, test.out, result, test.in)
		}
	}

}

func TestCountRepeatedCharacters(t *testing.T) {
	type Test struct {
		c1  map[string]int
		c2  map[string]int
		out int
	}

	var testsCases = []Test{
		{
			map[string]int{"a": 1},
			map[string]int{"a": 1},
			1,
		},
	}

	for i, test := range testsCases {
		result := countRepeatedCharacters(test.c1, test.c2)
		if !reflect.DeepEqual(result, test.out) {
			t.Errorf("Error in the test N°%v, Expected %v, but got %v. For the input %v and %v", i+1, test.out, result, test.c1, test.c2)
		}
	}
}
