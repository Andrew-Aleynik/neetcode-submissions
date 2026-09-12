func isPalindrome(s string) bool {
	isAlphanumeric := func(b byte) bool {
		return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9')
	}
	toLower := func(b byte) byte {
		if b >= 'A' && b <= 'Z' {
			return b + 32
		}
		return b
	}

	for i, j := 0, len(s)-1; i < j; {
		if !isAlphanumeric(s[i]) {
			i++
			continue
		}
		if !isAlphanumeric(s[j]) {
			j--
			continue
		}
		left := toLower(s[i])
		right := toLower(s[j])
		if left != right {
			return false
		}
		i++
		j--
	}
	return true
}

