package permutations

import (
	"fmt"
	arraymanipulation "myalgo/arrayManipulation"
	"strconv"
)

// Converts the given integer n in its factoradical form
func IToFactx(n int) (int, error) {
	var divisor int = 1
	var out string = ""
	for n >= 1 {
		toAdd := n % divisor
		out = strconv.Itoa(toAdd) + out
		n = n / divisor
		divisor++
	}
	res, err := strconv.Atoi(out)
	if err != nil {
		return -1, err
	}
	return res, nil
}

// Used internally in FactxToI
func factorial(n int) int {
	if n <= 0 {
		return 1
	}
	res := 1
	for n > 0 {
		res *= n
		n--
	}
	return res
}

// Converts the factoradical form f in an integer
func FactxToI(f int) int {
	strform := strconv.Itoa(f)
	fmt.Println("Strform", strform)
	res := 0
	for c := range strform {
		cifra := strform[c] - 48
		moltiplicatore := factorial(len(strform) - c - 1)
		res += int(cifra) * moltiplicatore
	}
	return res
}

// Calculate the nth permutation of a set using the factoradical system.
// Returns a string
func NthPermutation(n int, set []int) string {
	n = n - 1
	var out string
	for i := len(set) - 1; i >= 0; i-- {
		res := int(n / factorial(i))
		remainder := n % factorial(i)
		out += strconv.Itoa(set[res])
		set = arraymanipulation.DeleteIndex(res, set)
		n = remainder
	}
	return out
}
