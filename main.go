package main

import (
	"fmt"
	numbertheory "myalgo/numberTheory"
	"time"
)

func main() {
	a := time.Now()
	c := 1

	for {
		tn := c * (c + 1) / 2
		div := numbertheory.FindDivisor(tn)
		if len(div) > 1000 {
			fmt.Println(tn)
			break
		}
		c++
	}
	fmt.Println(time.Since(a))
}
