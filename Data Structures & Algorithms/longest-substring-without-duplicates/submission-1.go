func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	counts := make(map[byte]int)
	for i, j := 0, 0; j < len(s); {
		if counts[s[j]] >= 1 {
			counts[s[i]]--
			i++
		} else {
			counts[s[j]]++
			j++
			if j - i > maxLen {
				maxLen = j - i
			}
		}
	}
	return maxLen
}
