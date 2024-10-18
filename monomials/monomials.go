package monomials

import "fmt"

type Monom struct {
	Value float32
	Grade int	
}

func (m Monom) PPrint () {
	if m.Grade == 0 {
		fmt.Println(m.Value)
	}
	if m.Grade == 1 {
		fmt.Println("x^"+fmt.Sprint(m.Grade))
	}
	if m.Grade == -1 {
		
	}

}