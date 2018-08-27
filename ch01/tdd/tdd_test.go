package main

import "testing"

type MutiplyTC struct {
	a   int
	b   int
	res int
}

var multiplyTCs = []MutiplyTC{
	{1, 2, 2},
	{1, 3, 3},
	{4, 3, 12},
	{-4, 3, -12},
	{0, 3, 0},
	{1, 1, 1},
}

func TestMultiply(t *testing.T) {
	for _, tc := range multiplyTCs {
		if r := multiply(tc.a, tc.b); r != tc.res {
			t.Errorf("%d x %d should be %d, and it is %d\n", tc.a, tc.b, tc.res, r)
		}
	}
}
