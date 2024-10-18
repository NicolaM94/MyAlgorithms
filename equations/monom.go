/**Package containing the base class to create and instance a monom.
* A monom is an algebraic object in the form of a*x.
 */
package equations

import (
	"math"
)


type monom struct {
	value float32
	degree int
}

func MakeMonom (value float32, degree int) monom {
	return monom{value: value, degree: degree}
}

func (m monom) Add (n monom)

func (m monom) Elevate (deg int) monom {
	return MakeMonom(float32(math.Pow(float64(m.value),float64(deg))),m.degree*deg)
}