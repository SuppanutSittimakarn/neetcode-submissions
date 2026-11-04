func isValid(str string) bool {
	if len(str)%2 != 0 {
		return false
	}
	stack := []string{}
	mp := map[string]string{
		"}": "{",
		")": "(",
		"]": "[",
	}
	for _, s := range str {
		v, f := mp[string(s)]
		if !f {
			stack = append(stack, string(s))
		} else if len(stack) > 0 && v == stack[len(stack)-1] {
			stack = stack[0 : len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}