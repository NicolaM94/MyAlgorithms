package fibonacci

import "math"

// Returns the math implementation of the fibonacci sequence at position n
func FibonacciMath(n int) int {
	sigZero := math.Pow(1.618, float64(n))
	sigHat := math.Pow(-0.618, float64(n))
	result := (1 / math.Sqrt(5)) * (sigZero - sigHat)
	return int(result)
}

// Returns the n(th) position of the Fibonacci sequence in a recursive way
func FibonacciRec(n int) int {
	if n <= 2 {
		return 1
	}
	return FibonacciRec(n-1) + FibonacciRec(n-2)
}

// Returns the n(th) position of the Fibonacci sequence in a iterative way
func FibonacciIter(n int) int {
	a, b := 1, 1
	for i := 3; i <= n; i++ {
		c := a + b
		a = b
		b = c
	}
	return b
}
