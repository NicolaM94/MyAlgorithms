package ruffini

import "fmt"

// Base structure to parse equation terms
type Equation struct {
	Terms    []int
	Divisors []int
}

// Pretty prints the equation
func (e Equation) PPrint() {
	maxg := len(e.Terms) - 1
	out := ""
	for _, t := range e.Terms {
		if t > 0 {
			out += "+"
		} else if t == 0 {
			maxg--
			continue
		}
		out += fmt.Sprint(t) + "x^" + fmt.Sprint(maxg)
		maxg--
	}
	fmt.Println(out[:len(out)-3])
}

// Finds all the divisors of the last term of e and stores them in e.Divisors.
// The method calculates both positive and negative values directly, since it's common that solutions are negative if non-positive and vice-versa
func (e *Equation) FindDivisors() {
	dividend := e.Terms[len(e.Terms)-1]
	e.Divisors = append(e.Divisors, 1, -1, dividend, dividend*(-1))
	for m := 2; m <= int(dividend/2); m++ {
		if m == e.Divisors[len(e.Divisors)-2] {
			break
		}
		if dividend%m == 0 {
			e.Divisors = append(e.Divisors, m, m*(-1), dividend/m, (-1)*(dividend/m))
		}
	}
}

// Applies Ruffini's rule and returns an array with the downgraded equation and the divisor applied with its sign inversed
func (e *Equation) RuffinisRule() ([]int, int) {
	e.FindDivisors()
	for _, d := range e.Divisors {
		tS := []int{e.Terms[0]}
		for _, t := range e.Terms[1:] {
			tS = append(tS, d*tS[len(tS)-1]+t)
		}
		if tS[len(tS)-1] == 0 {
			return tS[0 : len(tS)-1], d * (-1)
		}
	}
	return nil, 0
}

/*
Code sample to reduce the given polynomial to its minimum grade form

func main() {

	e := Equation{
		Terms: []int{1, 4, 1, -6}}
	e.FindDivisors()
	out, div := e.RuffinisRule()
	for out != nil {
		fmt.Println("Solution: ", out, "Divisor: ", div)
		e = Equation{Terms: out}
		e.FindDivisors()
		out, div = e.RuffinisRule()
	}
}

# Solution:  [1 5 6] Divisor:  -1
# Solution:  [1 3] Divisor:  2
# Solution:  [1] Divisor:  3


*/
