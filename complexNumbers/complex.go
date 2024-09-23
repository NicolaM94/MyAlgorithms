package complexnumbers

import (
	"fmt"
	"math"
)

type Cmpx struct {
	Real float64
	Imag float64
}

// Pretty prints the complex number cn in its algebraic form.
func (cn *Cmpx) PrettyAlgebraic() {
	if cn.Imag == 0 && cn.Real == 0 {
		fmt.Println(0)
		return
	}
	var out string
	// Parse real part
	if (cn.Real > 0) {
		out += "+"
		out += fmt.Sprint(cn.Real)
	} else if (cn.Real < 0) {
		out += fmt.Sprint(cn.Real)
	}
	// Parse the imaginary part
	if (cn.Imag > 0) {
		out += "+"
		out += fmt.Sprint(cn.Imag)
		out += "i"
	} else {
		out += fmt.Sprint(cn.Imag)
		out += "i"
	}
	fmt.Println(out)
}

// Calculates the module of the complex number.
func (cn *Cmpx) Module() float64 {
	var base = cn.Real * cn.Real + cn.Imag * cn.Imag
	return math.Sqrt(base)
}

// Calculates the anomaly (angle) of the complex number.
func (cn *Cmpx) Anomaly() float64 {
	if (cn.Real == 0) {
		if (cn.Imag > 0) {
			return math.Pi/2
		} else if (cn.Imag < 0) {
			return math.Pi * 3/2
		}
	}
	return math.Atan2(cn.Imag,cn.Real)
}

// Returns the conjugate of the complex number: if z = a + bi it returns a - bi.
func (cn *Cmpx) Conjugate() Cmpx {
	return Cmpx{cn.Real, cn.Imag * -1}
}

// Pretty prints the complex number cn in its trig form.
func (cn *Cmpx) PrettyTrigonometric () {
	var out string
	out += fmt.Sprint(cn.Module())
	out += "(cos "
	out += fmt.Sprint(cn.Anomaly())
	out += " + sin "
	out += fmt.Sprint(cn.Anomaly())
	out += ")"
	fmt.Println(out)
}
