func maxArea(heights []int) int {
	res := 0

	for l, r := 0, len(heights)-1; l < r; {
		water := (r - l) * min(heights[l], heights[r])
		res = max(res, water)

		if heights[l] <= heights[r] {
			l++
		} else if heights[l] > heights[r] {
			r--
		}
	}

	return res
}