package sorting

import "fmt"

// FIXME: Here
func QuickSort (array []int, pivot int) [][]int {
	i := 0
	j := len(array)-1
	counter := 1
	for i < j {
		fmt.Println("loop ", counter)
		m := i
		for array[m] < pivot {
			m++
		}
		fmt.Println("Array[m]: ",array[m])
		n := j
		for array[n] > pivot {
			n--
		}
		fmt.Println("Array[n]: ",array[n])
		if m >= n {
			i = m
			j = n
			break
		}
		temp := array[n]
		array[n] = array[m]
		array[m] = temp
		i++
		j--
		fmt.Println("New array: ",array)
		fmt.Println(" ")
		counter++
	}
	return [][]int{array[:i],array[j:]}
}