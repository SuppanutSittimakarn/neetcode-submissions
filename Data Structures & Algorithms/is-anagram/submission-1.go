func isAnagram(str1 string, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	for _, s1 := range str1 {
		tmp := len(str2)
		for _, s2 := range str2 {
			if s1 == s2 {
				str2 = strings.Replace(str2, string(s1), "", 1)
				break
			}
		}

		if len(str2) == tmp {
			return false
		}
	}
	return true
}