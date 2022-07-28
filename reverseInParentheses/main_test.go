package main

import "testing"

type Test struct {
	in  string
	out string
}

var testsCases = []Test{
	{
		"(bar)",
		"rab",
	},
	{
		"a(bc)d",
		"acbd",
	},
	{
		"a(bc)",
		"acb",
	},
	{
		"foo(bar)baz(blim)",
		"foorabbazmilb",
	},
}

func TestLazyCutter(t *testing.T) {

	for i, test := range testsCases {
		result := reverseInParentheses(test.in)
		if result != test.out {
			t.Errorf("Error in the test N°%v, Expected %v, but got %v. For the input %v", i+1, test.out, result, test.in)
		}
	}

}
