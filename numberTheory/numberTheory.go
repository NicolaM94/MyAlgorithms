package numbertheory

import "math"

func CheckTriangular(n int) bool {
	base := float64(1 + 8*n)
	sqroot := math.Sqrt(base)
	return sqroot/2-math.Floor(sqroot/2) == 0.5
}

func CreateTriangular(nth int) int {
	return nth * (nth + 1) / 2
}

func FindDivisors(n int) (out []int) {
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			out = append(out, i)
		}
	}
	return
}
