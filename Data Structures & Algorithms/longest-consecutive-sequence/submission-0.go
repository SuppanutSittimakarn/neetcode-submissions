
func longestConsecutive(nums []int) int {
	set := make(map[int]struct{})
	for _, n := range nums {
		set[n] = struct{}{}
	}

	longest := 0
	for _, n := range nums {
		if _, f := set[n-1]; !f {
			length := 1
			for i := n; ; i++ {
				if _, f := set[i+1]; !f {
					break
				}
				length++
			}
			if length > longest {
				longest = length
			}
		}
	}

	return longest
}
