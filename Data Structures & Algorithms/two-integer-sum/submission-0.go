func twoSum(nums []int, target int) []int {
	r := make(map[int]int)

	for i, n := range nums {
		if a, f := r[target-n]; f {
			return []int{a, i}
		} else {
			r[n] = i
		}
	}

	return []int{}
}