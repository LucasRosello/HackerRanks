package main

import "testing"

type Test struct {
	in  int
	out int
}

var testsCases = []Test{
	{
		1,
		2,
	},
	{
		2,
		4,
	},
	{
		3,
		7,
	},
	{
		4,
		11,
	},
	{
		42,
		904,
	},
}

func TestLazyCutter(t *testing.T) {

	for i, test := range testsCases {
		result := OptimizedLazyCutter(test.in)
		if result != test.out {
			t.Errorf("Error in the test N°%v, Expected %v, but got %v. For the input %v", i+1, test.out, result, test.in)
		}
	}

}

func BenchmarkSimpleLazyCutter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SimpleLazyCutter(100000)
	}
}

func BenchmarkOptimizedLazyCutter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OptimizedLazyCutter(100000)
	}
}
