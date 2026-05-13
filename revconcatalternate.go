// RevConcatAlternate receives two slices and returns a new slice
// with alternated values in reverse order based on length priority.
func RevConcatAlternate(slice1, slice2 []int) []int {
	len1 := len(slice1)
	len2 := len(slice2)
	totalLen := len1 + len2
	result := make([]int, totalLen)

	// i and j represent the current index (pointer) for each slice
	i, j := len1-1, len2-1

	for k := 0; k < totalLen; k++ {
		// Logic: Pick from slice1 if it's currently "longer" (more elements left),
		// or if lengths are equal, or if slice2 is exhausted.
		if i >= 0 && (i >= j || j < 0) {
			result[k] = slice1[i]
			i--
		} else if j >= 0 {
			result[k] = slice2[j]
			j--
		}
	}
	return result
}
