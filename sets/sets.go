package sets

// Module containing all the functions referred to set theory.
// Since the modules takes a functional aproach, a list of int is to be considered as a set, meaning that all the functions refers to lists of integers as sets.
// Notice that sets are considered to contain unique elements (no repetitions)

// Checks if set contains element
// Not exported, used only as module internal
func contains(element int, set []int) bool {
	for s := range set {
		if element == set[s] {
			return true
		}
	}
	return false
}

// Returns true if A is empty
func IsEmpty (A []int) bool {
	return len(A) == 0
}

// Returns true if B is a subset of A
func IsSubset (B,A []int) bool {
	for b := range B {
		if !contains(B[b],A) {
			return false
		}
	}
	return true
}

// Returns true if B is a proper subset of A, meaning that A contains other elements besides B's.
func IsProperSubset (B,A []int) bool {
	if len(B) != len(A) && IsSubset(B,A) {
		return true
	}
	return false
}

//TODO: Da fare qui, l'insieme delle parti è complesso perchè vanno scritte le combinazioni di n -k elementi, per k da 1 a n-1.
func PartsSetOf(A []int) (p [][]int) {
	return
}

// Returns the intersection set between A and B
// Notice that the function only checks if values of A are present in B, thanks the the commutative property of intersection.
func Intersection (A,B []int) (intersection []int) {
	for a := range A {
		if contains(A[a], B) {
			intersection = append(intersection, A[a])
		}
	}
	return intersection
}

// Returns true if C is a valid intersection of A and B, meaning that C elements are contained in A AND B.
func IsIntersection (C, A, B []int) bool {
	for c := range C {
		if !contains(C[c], A) || !contains(C[c],B) {
			return false
		}
	}
	return true
}

// Returns true if A and B are disjoned, meaning they have no elements in common.
func AreDisjoined(A,B []int) bool {
	return len(Intersection(A,B)) == 0
}


// Returns the union set between A and B.
// Notice that the function only checks if values of B are present in A, since the union is commutative
func Union (A,B []int) (union []int) {
	union = A
	for b := range B {
		if !contains(B[b], A) {
			union = append(union, B[b])
		}
	}
	return
}

// Returns true if C is a valid union of A and B, meaning that C elements are contained in A OR B.
func IsUnion (C, A, B []int) bool {
	if len(C) < len(A)+len(B) {
		return false
	}
	for c := range C {
		if !contains(C[c], A) && !contains(C[c],B) {
			return false
		}
	}
	return true
}

// Difference

// Returns the difference between A and B, as A\B.
// Notice that the difference does not commute, meaning that A\B != B\A
func Difference(A, B []int) (difference []int) {
	for a := range A {
		if !contains(A[a], B) {
			difference = append(difference, A[a])
		}
	}
	return
}

// Return the complementary set of B in regards of A. This is possible only if B is a proper subset of A and the
// complementary set will be the difference A\B
func ComplementarySet(B,A []int) []int {
	if IsProperSubset(B,A) {
		return Difference(A,B)
	}
	return []int{}
}

// Return the symmetric difference between A and B, aka the difference between the union and the intersection of A and B.
// The result is the array of the elements of A not in B and the elements of B not in A
func SymmetricDifference (A,B []int) []int {
	return Difference(Union(A,B),Intersection(A,B))
}


// Returns A x B.
// The product does not commute (AxB != BxA)
func Product(A,B []int) (product [][]int) {
	for a := range A {
		for b:= range B {
			product = append(product, []int{A[a],B[b]})
		}
	}
	return
}

// Returns the square of A, meaning AxA.
// An example is R^2
func Square(A []int) (square [][]int) {
	for a := range A {
		for b := range A {
			square = append(square, []int{A[a],A[b]})
		}
	}
	return square
}

// Returns the diagonal of A^2
func DiagonalSquareOf(A []int) (diagonal [][]int) {
	for a := range A {
		diagonal = append(diagonal, []int{A[a],A[a]})
	}
	return
}