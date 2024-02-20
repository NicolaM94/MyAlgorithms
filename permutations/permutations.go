package permutations

import "fmt"

func In(test int, array []int) bool {
	for i := range array {
		if array[i] == test {
			return true
		}
	}
	return false
}

func NextPermutation() int {
	result := 0
	found := []int{}
	for i := 1; i <= 999; i++ {
		for j := 1; j <= 999; j++ {

			if In(i*j, found) {
				continue
			}

			stringed := fmt.Sprint(i) + fmt.Sprint(j) + fmt.Sprint(i*j)
			if len(stringed) != 9 {
				continue
			}

			counter := make(map[int]int)
			for n := 1; n <= 9; n++ {
				counter[n] = 0
			}
			for _, f := range stringed {
				counter[int(f-48)]++
			}
			for _, k := range counter {
				if counter[k] != 1 {
					continue
				}
			}

			found = append(found, i*j)
			result += i * j

		}
	}
	return result
}
