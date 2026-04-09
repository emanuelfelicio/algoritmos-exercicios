package algoritmosexercicios

func isValid(s string) bool {
	m := map[rune]rune{
		']': '[',
		')': '(',
		'}': '{',
	}
	stack := make([]rune, 0)
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stack = append(stack, v)
		} else {
			
			if len(stack) > 0 && stack[len(stack)-1] == m[v] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	if len(stack) > 0 {
		return false
	}
	//([{}])
	return true
}
