package complexnumbers

import "math"

// Returns the sum of the two complex numbers z1 and z2.
func AddComplex(z1, z2 Cmpx) Cmpx {
	var out Cmpx
	out.Real = z1.Real + z2.Real
	out.Imag = z1.Imag + z2.Imag
	return out
}

// Returns the subtraction of the two complex numbers z1 and z2.
func SubtractComplex(z1, z2 Cmpx) Cmpx {
	var out Cmpx
	out.Real = z1.Real - z2.Real
	out.Imag = z1.Imag - z2.Imag
	return out
}

// Returns the product between the two complex numbers z1 and z2.
func ProductComplex(z1, z2 Cmpx) Cmpx {
	var out Cmpx
	out.Real = (z1.Real * z2.Real) - (z1.Imag * z2.Imag)
	out.Imag = (z1.Real * z2.Imag) + (z1.Imag * z2.Real)
	return out
}

// Returns the division between two complex numbers z1 and z2.
func DivideComplex(z1, z2 Cmpx) Cmpx {
	var out Cmpx
	var divisor float64 = z2.Real * z2.Real + z2.Imag * z2.Imag
	var numOne = z1.Real * z2.Real + z1.Imag * z2.Imag
	var numTwo = z1.Imag * z2.Real - z1.Real * z2.Imag
	out.Real = numOne/divisor
	out.Imag = numTwo/divisor
	return out
}

// Returns the nth power of z
func PowerComplex(z Cmpx, n int) Cmpx {
	var out = z
	for n-1 > 0 {
		out = ProductComplex(out, z)
		n--
	}
	return out
}

// Returns the nth roots of z.
// Notice that the nth root of z has exactly n solution, calculated by a factor k={0,1,...,n-1}
func RootComplex(z Cmpx, n int) []Cmpx {
	var m = z.Module()
	var a = z.Anomaly()

	var out []Cmpx
	var solution Cmpx

	for n-1 >= 0 {

		var cos = math.Cos((a*2*float64(n-1)*math.Pi/float64(n)))
		var sin = math.Sin((a*2*float64(n-1)*math.Pi/float64(n)))

		solution.Real = math.Sqrt(m) * cos
		solution.Imag = math.Sqrt(m) * sin

		out = append(out, solution)
		n--
	}
	return out
}