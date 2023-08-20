package arraymanipulation

import "fmt"

func Insert(element, index int, array []int) []int {
	out := []int{}
	out = append(out, array[:index]...)
	out = append(out, element)
	fmt.Println("Array", array)
	out = append(out, array[index:]...)
	return out
}

// Shifts an array left to delete the first item
func DeleteLeft(in []int) []int {
	for e := 1; e < len(in); e++ {
		in[e-1] = in[e]
	}
	return in[:len(in)-1]
}

// Shift an array right to delete the last item
func DeleteRight(in []int) []int {
	for e := len(in) - 1; e >= 1; e-- {
		in[e] = in[e-1]
	}
	return in[1:]
}

// Delete the element at index i from the array
func DeleteIndex[T comparable](i int, array []T) []T {
	out := []T{}
	for n := 0; n <= i-1; n++ {
		out = append(out, array[n])
	}
	for n := i + 1; n < len(array); n++ {
		out = append(out, array[n])
	}
	return out
}

// Deletes the first occurence of the element found in the array
func DeleteElem[T comparable](s []T, k T) []T {
	for index := 0; index < len(s); index++ {
		if s[index] == k {
			return DeleteIndex(index, s)
		}
	}
	return s
}

// Shifts all the array values to one right and then inserts toInsert at array[0].
// Equivalent to delete the last, inserting a new one at the start.
func ShiftRightInsertLeft(array []int, toInsert int) []int {

	for i := len(array) - 2; i >= 0; i-- {
		array[i+1] = array[i]
	}
	array[0] = toInsert
	return array
}

// Shifts all the array values to one left and the inserts toInsert at the end of the array
// Equivalent to delete the first, inserting a new one at the end.
func ShiftLeftInsertRight(array []int, toInsert int) []int {

	for i := 1; i <= len(array)-1; i++ {
		array[i-1] = array[i]
	}
	array[len(array)-1] = toInsert
	return array

}
