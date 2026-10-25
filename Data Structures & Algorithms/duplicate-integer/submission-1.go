func hasDuplicate(nums []int) bool {
	seen := make(map[int]bool)

	for _, n := range nums {
		_, found := seen[n]
		if found {
			return found
		} else {
			seen[n] = true
		}

	}
	return false
}