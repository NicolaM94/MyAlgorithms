package numbertheory

import "math"

// Checks if a number is a triangular number
func CheckTriangular(n int) bool {
	base := float64(1 + 8*n)
	sqroot := math.Sqrt(base)
	return sqroot/2-math.Floor(sqroot/2) == 0.5
}

// Create the nth triangular number
func CreateTriangular(nth int) int {
	return nth * (nth + 1) / 2
}

// Finds divisors of the number n
func FindDivisor(n int) (out []int) {
	out = append(out, 1)
	out = append(out, n)

	for j := 2; j < n/2; j++ {
		if j == out[len(out)-2] || j == out[len(out)-1] {
			break
		}
		if j*j == n {
			out = append(out, j)
			continue
		}
		if n%j == 0 {
			out = append(out, j)
			out = append(out, n/j)
		}
	}
	return
}
