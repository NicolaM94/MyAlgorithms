package sorting

import (
	
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
	}
	return array
}

// Insertion sort algorithm
// TODO: Fix insertion
func InsertionSort (array []int) []int {
	return array
}

// Bubble sort algorithm
//TODO: fix bubble
func BubbleSort (array []int) []int {
	for i := 0; i <= len(array)-2; i++ {
		scambi := false
		for j := 1 ; j<= len(array)-1-i;j++ {
			if array[j-1] > array[j] {
				temp := array[j-1]
				array[j-1] = array[j]
				array[j] = temp
				scambi = true
			}
			if scambi {
				break
			}
		}
	}
	return array
}