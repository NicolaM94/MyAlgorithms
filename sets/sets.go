package sets



// Private function to check if element x is contained by A
// This function is only used in this modules
func contains (x int, A []int) bool {
	for a := range A {
		if A[a] == x {
			return true
		}
	}
	return false
}
// Returns true if B is a subset of A
func IsIncluded (A, B []int) bool {
	if len(B) == 0 {
		return true
	}
	for i := range B {
		if !contains(B[i], A) {
			return false
		}
	}
	return true
}

// Returns true if B is a proper subset of A
func IsProperlyIncluded (A, B []int) bool {
	if len(B) == 0 {
		return false
	}
	collector := []int{}
	for b := range B {
		if !contains(B[b], A) {
			return false
		}
		collector = append(collector, B[b])
	}
	if len(collector) < len(A) {
		return true
	}
	return false
}

// Returns the set of the parts of A
func SetOfTheParts (A []int) (out [][]int) {
	if len(A) == 0 {
		return
	}
	out = append(out, []int{})
	out = append(out, A)
	for i := 0; i < len(A)-1; i++ {
		for j := i+1; j < len(A); j++ {
			out = append(out, []int{A[i],A[j]})
		}
	}
	return
}

// Returns the intersection set between A and B
func Intersection (A, B []int) (out []int) {
	for a := range A {
		if contains(A[a], B) {
			out = append(out, A[a])
		}
	}
	return
}

// Returns true if A and B are disjointed
func AreDisjointed (A, B []int) bool {
	return len(Intersection(A,B)) == 0
}

// Returns the union between A and B
func Union (A, B []int) (out []int) {
	out = append(out, A...)
	out = append(out, B...)
	return
}

// Returns the set of the elements contained by A but not by B
func Difference (A, B []int) (out []int) {
	for a := range A {
		if !contains(A[a],B) {
			out = append(out, A[a])
		}
	}
	return out
}

// Returns the simmetric difference between A and B
func SimmetricDifference (A, B []int) []int {
	union := Union(A,B)
	intersection := Intersection(A,B)
	return Difference(union,intersection)
}

// Returns the cartesian product of AxB
func CartesianProduct (A, B []int) (out [][]int) {
	for a := range A {
		for b := range B {
			out = append(out, []int{A[a],B[b]})
		}
	}
	return
}

// Returns true when B is a cover of A, meaning that each set of B contains at least one element of A;
// This function is resource intensive, do not use it with high volumes sets.
func IsACover(A []int, B [][]int) bool {
	for a := range A {
		tempContains := false
		for b := range B {
			tempContains = contains(A[a], B[b])
			if tempContains {
				break
			}
		}
		if tempContains {
			continue
		}
		return false
	}
	return true
}

// Returns true when B is a partition of A, meaning that each set of B contains at least one element of A;
// This function is resource intensive, do not use it with hight volumes sets.
func IsAPartition (A []int, B [][]int) bool {
	for i := 0; i < len(B)-1; i++ {
		for j := i+1; j < len(B); j++ {
			if len(Intersection(B[i],B[j])) != 0 {
				return false
			}
		}
	}
	return IsACover(A,B)
}