package arithmetic

import "testing"

type tc struct {
	ops []int
	res int
}

func TestSum(t *testing.T) {

	tcs := []tc{
		{[]int{1, 2, 3}, 6},
		{[]int{6, -3, 3}, 6},
	}

	for _, tc := range tcs {
		res := Sum(tc.ops...)
		if tc.res != res {
			t.Errorf("Sum(%d) is %d, expected %d\n", tc.ops, res, tc.res)
		}
	}

}

func TestSubstract(t *testing.T) {

	tcs := []tc{
		{[]int{1, 2, 3}, -4},
		{[]int{6, -3, 3}, 6},
	}

	for _, tc := range tcs {
		res := Substract(tc.ops...)
		if tc.res != res {
			t.Errorf("Substract(%d) is %d, expected %d\n", tc.ops, res, tc.res)
		}
	}

}

func TestMultiply(t *testing.T) {

	tcs := []tc{
		{[]int{1, 2, 3}, 6},
		{[]int{6, -3, 3}, -54},
	}

	for _, tc := range tcs {
		res := Multiply(tc.ops...)
		if tc.res != res {
			t.Errorf("Multiply(%d) is %d, expected %d\n", tc.ops, res, tc.res)
		}
	}

}
