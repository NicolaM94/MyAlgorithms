package sorting

import (
	"fmt"
)

// Selection sort algorithm
func SelectionSort (array []int) []int {
	for i:=0;i<=len(array)-2;i++ {
		minValue := i+1
		for j:=i+2;j<len(array);j++ {
			if array[j] < array[minValue] {
				minValue = j
			}
		}
		if array[i] > array[minValue] {
			temp := array[i]
			array[i] = array[minValue]
			array[minValue] = temp
		}
		
		fmt.Println(array)
		
	}
	return array
}

// Insertion sort algorithm
func InsertionSort (array []int) []int {
	for i:=1;i<=len(array)-1;i++ {
		for j := i-1;j==0;j-- {
			if array[i] < array[j] {
				temp := array[j]
				array[j] = array[i]
				array[i] = temp
			}
		}
	}
	return array
}