package arraymanipulation

// Shifts all the array values to one right and then inserts toInsert at array[0]
func ShiftRightToInsert (array []int, toInsert int) []int {

	for i:=len(array)-2;i>=0;i-- {
		array[i+1] = array[i]
	}
	array[0] = toInsert
	return array
}

// Shifts all the array values to one left and the inserts toInsert at the end of the array
func ShiftLeftToInsert (array []int, toInsert int) []int {

	for i:= 1;i<=len(array)-1;i++ {
		array[i-1] = array[i]
	}
	array[len(array)-1] = toInsert
	return array

}