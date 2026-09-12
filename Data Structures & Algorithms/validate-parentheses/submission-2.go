func isValid(str string) bool {
    stack := []rune{}
	for _, s := range str {
		if s == '(' || s == '[' || s == '{' {
			stack = append(stack, s)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if top == '(' && s != ')' || top == '[' && s != ']' || top == '{' && s != '}' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
