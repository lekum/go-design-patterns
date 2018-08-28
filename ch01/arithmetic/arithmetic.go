package arithmetic

func Sum(args ...int) int {
	res := 0
	for _, v := range args {
		res += v
	}
	return res
}

func Substract(args ...int) int {
	res := args[0]
	for _, v := range args[1:] {
		res -= v
	}
	return res
}

func Multiply(args ...int) int {
	res := args[0]
	for _, v := range args[1:] {
		res *= v
	}
	return res
}
