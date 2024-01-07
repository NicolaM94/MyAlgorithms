package sets

import "fmt"

func containsArray (A []int, B [][]int) bool {
	out := false
	LOOP:
		for _, b := range B {
			for _, a := range A {
				if !contains(a,b) {
					continue LOOP
				}
			}
			out = true
			break
		}
	return out
}



func IsTopology (X []int, T [][]int) {
	fmt.Println(containsArray(X, T))
}