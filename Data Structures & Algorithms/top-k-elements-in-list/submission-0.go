
func topKFrequent(nums []int, k int) []int {
	r := []int{}

	count := make(map[int]int)
	freq := make([][]int, len(nums)+1)
	for _, n := range nums {
		count[n]++
	}

	for a, n := range count {
		freq[n] = append(freq[n], a)
	}

	for i := len(freq) - 1; i < len(freq); i-- {
		for _, num := range freq[i] {
			r = append(r, num)
			if len(r) == k {
				return r
			}
		}
	}

	return r
}
