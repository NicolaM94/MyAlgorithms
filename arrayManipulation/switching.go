package arraymanipulation

// Switches two elements at indices a and b of an array 
func SwitchIndexes (a,b int, array []int) []int {
	temp := array[a]
	array[a] = array[b]
	array[b] = temp
	return array
}

// Switches two elements a and b (first occurrencies from the left) of an array
func SwitchElements (a,b int, array []int) []int {
	aInd := 0
	bInd := 0
	for i,k := range array {
		if k == a {
			aInd = i
		}
	}
	for i,j := range array {
		if j == b {
			bInd = i
		}
	}
	temp := array[aInd]
	array[aInd] = array[bInd]
	array[bInd] = temp
	return array
}

// Reverse an array in place
func Reverse (a []int) []int {
	for n:=0;n<=len(a)/2;n++ {
		temp := a[n]
		a[n] = a[len(a)-n-1]
		a[len(a)-n-1] = temp
	}
	return a
}