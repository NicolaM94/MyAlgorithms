package equations

import (
	"myalgo/sommations"
	"myalgo/str"
	"strconv"
)

// Parse the basic monomial a*x
// Returns (a, true) if x != 0
// Returns (a, false) else
func baseParser(monom string) (int, bool) {
	if monom[len(monom)-1] == 'x' {
		if len(monom) == 1 {
			monom = str.Insert("1", monom, 0)
		}
		if len(monom) == 2 {
			monom = str.Insert("1", monom, 1)
		}
		a, err := strconv.Atoi(monom[:len(monom)-1])
		if err != nil {
			panic(err)
		}
		return a, true
	} else {
		a, err := strconv.Atoi(monom)
		if err != nil {
			panic(err)
		}
		return a, false
	}
}

// Solves linear (first grade) equations
func LinearSolver(equation string) float32 {
	xs := []int{}
	ns := []int{}
	out := string(equation[0])
	for _, c := range equation[1:] {
		// When finds a sign
		if c == '+' || c == '-' {
			// Stops and parse out
			// If it's x monomial
			value, isx := baseParser(out)
			if isx {
				xs = append(xs, value)
			} else {
				ns = append(ns, value)
			}
			out = ""
		}
		// Adds last carachter to out
		out += string(c)
	}
	// Parse the remaining part of out
	value, isx := baseParser(out)
	if isx {
		xs = append(xs, value)
	} else {
		ns = append(ns, value)
	}
	result := -1 * float32(sommations.IntArraySum(ns)) / float32(sommations.IntArraySum(xs))
	return result

}
