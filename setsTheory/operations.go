package setstheory

// Calculates the intersection between two sets, as the logic gate operation AND
func Intersection(setA, setB []int) []int {
	union := []int{}
	for a := range setA {
		for b := range setB {
			if setA[a] == setB[b] {
				union = append(union, setA[a])
			}
		}
	}
	return union
}

// Calculates the union between two sets, as the logic gate operation OR
func Union(setA, setB []int) []int {
	out := []int{}
	out = append(out, setA...)
	out = append(out, setB...)
	return out
}

// Calculates the difference between two sets, as the logic gate operation XOR
func Difference(setA, setB []int) []int {
	out := []int{}
LoopOfA:
	for a := range setA {
		for b := range setB {
			if setA[a] == setB[b] {
				continue LoopOfA
			}
		}
		out = append(out, setA[a])
	}
	return out
}
