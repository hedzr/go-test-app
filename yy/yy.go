package yy

func Factorial(num int) int {
	return factorialTailRecursive(num)
}

func factorialTailRecursive(num int) int {
	return factorial(1, num)
}

func factorial(accumulator, val int) int {
	if val <= 1 {
		return accumulator
	}
	return factorial(accumulator*val, val-1)
}
