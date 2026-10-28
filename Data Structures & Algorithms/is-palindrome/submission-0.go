func isPalindrome(str string) bool {
	re := regexp.MustCompile(`[^a-zA-Z0-9]`)
	str = strings.ToLower(re.ReplaceAllString(str, ""))
	for i := range str {
		if i > len(str)/2 {
			return true
		}
		a := str[i]
		b := str[len(str)-1-i]
		if a != b {
			return false
		}
	}

	return true
}