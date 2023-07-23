package orderingalgorithms

import (
	"fmt"
)

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