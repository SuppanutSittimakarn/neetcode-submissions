func longestConsecutive(nums []int) int {
	longest := 0

	mp := make(map[int]int)
	for _, n := range nums {
		if mp[n] == 0 {
			left := mp[n-1]
			right := mp[n+1]

			sum := left + right + 1
			mp[n] = sum

			mp[n-left] = sum
			mp[n+right] = sum

			if sum > longest {
				longest = sum
			}
		}
	}

	return longest
}