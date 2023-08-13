package arraymanipulation

// Compares arrays a and b. Returns true if they are equal (considering elements order)
func CompareOrder (a,b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for x := range a {
		if a[x] != b[x] {
			return false
		}
	}
	return true
}

// Checks if int a is present in array b
func In (a int, b []int) bool {
	for x:= range b {
		if b[x] == a {
			return true
		}
	}
	return false
}

// Compares arrays a and b. Returns true if they are equal (ignoring elements order)
func CompareUnorder (a,b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for x := range a {
		if !In(a[x],b) {
			return false
		}
	}
	return true
}

// Returns the union of sets a and b
func Union (a,b []int) []int {
	for x := range a {
		if !In(a[x],b) {
			b = append(b, a[x])
		}
	}
	return b
}

// Returns the intersection of sets a and b
func Intersection (a,b []int) []int {
	out := []int{}
	for x := range a {
		if In(a[x],b) {
			out = append(out, a[x])
		}
	}
	return out
}