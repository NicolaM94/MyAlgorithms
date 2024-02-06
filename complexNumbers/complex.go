package complexnumbers

import (
	"fmt"
	"math"
)

// WIP library to work with complex numbers, works ONLY with n being natural (N)

// Base structure type for the complex number
// Accepts only the real and imaginary part and + or - sign
type ComplexN struct {
	RealPart float32	
	ImagPart float32
}

// Calculate the modulo of the complex number
func (c ComplexN) Modulo () (int) {
	return int(c.RealPart*c.RealPart + c.ImagPart * c.ImagPart)
}

// Calculates the conjuigate of the complex number
func (c ComplexN) Conjugate () ComplexN {
	return ComplexN{c.RealPart, c.ImagPart * -1}
}

// Add to c another complex number n not in place, returning a new value.
// Used also for subtraction.
func (c ComplexN) Add (n ComplexN) ComplexN {
	return ComplexN{
		RealPart: c.RealPart + n.RealPart,
		ImagPart: c.ImagPart + n.ImagPart,
	}
}

// Multiplication
func (c ComplexN) Product (n ComplexN)  ComplexN {
	return ComplexN{
		RealPart: (c.RealPart*n.RealPart)-(c.ImagPart*n.ImagPart),
		ImagPart: (c.RealPart*n.ImagPart)+(c.ImagPart*n.RealPart),
	}
}


func (c ComplexN) Division (n ComplexN) ComplexN {
	return ComplexN{
		RealPart: c.Product(n.Conjugate()).RealPart / float32(n.Modulo()),
		ImagPart: c.Product(n.Conjugate()).ImagPart / float32(n.Modulo()),
	}
}


// Print the number c in its algebraic form
// Makes use of fmt.Sprint function to convert integers to string
func (c ComplexN) PrintAlgebraic () {
	out := fmt.Sprint(c.RealPart)
	if c.ImagPart >= 0 {
		out = out + "+"
	}
	out = out + fmt.Sprint(c.ImagPart)
	out = out + "i"
	print(out)
}

func (c ComplexN) PrintTrigonometric () {

	angle := math.Atan(math.Tan(float64(c.ImagPart)/float64(c.RealPart)))

	print(fmt.Sprintf("%v(cos %vπ + isin %vπ)",c.Modulo(), angle,angle))

}