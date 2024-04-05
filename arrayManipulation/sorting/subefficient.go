/*
Package containing array ordering algorithms which sort arrays in a sub-efficient way
*/
package sorting

// Selection sort algorithm
func SelectionSort(array []string) []string {
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
func BubbleSort(array []byte) []byte {
	wall := len(array)
	for wall > 0 {
		for f := 0; f <= wall-2; f++ {
			for g := f + 1; g <= wall-1; g++ {
				if array[f] > array[g] {
					temp := array[g]
					array[g] = array[f]
					array[f] = temp
				}
			}
		}
		wall--
	}
	return array
}
