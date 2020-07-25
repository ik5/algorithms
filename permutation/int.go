package permutation

// IntSlice returns a list of all permutations for n times.
func IntSlice(n int, data []int) [][]int {
	result := [][]int{}

	if n < 1 {
		return result
	}

	if n == 1 {
		return append(result, data)
	}

	return result
}
