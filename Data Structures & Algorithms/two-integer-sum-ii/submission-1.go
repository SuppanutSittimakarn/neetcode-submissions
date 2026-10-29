func twoSum(numbers []int, target int) []int {
	for l, r := 0, len(numbers)-1; l < r; {
		sum := numbers[l] + numbers[r]
		if sum > target {
			r -= 1
		} else if sum < target {
			l += 1
		} else {
			return []int{l + 1, r + 1}
		}
	}

	return []int{}
}