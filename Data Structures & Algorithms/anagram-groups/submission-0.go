func groupAnagrams(strs []string) [][]string {
	res := make(map[[26]int][]string)

	for _, str := range strs {
		var count [26]int
		for _, s := range str {
			count[s-'a']++
		}

		res[count] = append(res[count], str)
	}

	answer := [][]string{}
	for _, r := range res {
		answer = append(answer, r)
	}

	return answer
}