/*
Package containing array ordering algorithms which sort arrays in a sub-efficient way
*/
package sorting

// Selection sort algorithm
func SelectionSort(array []int) []int {
	for i := 0; i <= len(array)-2; i++ {
		minValue := i + 1
		for j := i + 2; j < len(array); j++ {
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
func InsertionSort(array []int) []int {
	for k := 1; k <= len(array)-1; k++ {
		for j := k - 1; j >= 0; j-- {
			if array[j] > array[j+1] {
				temp := array[j+1]
				array[j+1] = array[j]
				array[j] = temp
			} else {
				continue
			}
		}
	}
	return array
}

// Bubble sort algorithm
func BubbleSort (array []int) []int {
	for i:=0;i<=len(array)-1;i++ {
		scambi := false
		for j:=1;j<=len(array)-2-i;j++ {
			if array[j-1] > array[j] {
				temp := array[j-1]
				array[j-1] = array[j]
				array[j] = temp
				scambi = true
			}
		}
		if !sorted {
			break
		}
	}
	return array
}
