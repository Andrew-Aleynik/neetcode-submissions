func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	locations := make(map[byte]int)
	for i, j := 0, 0; j < len(s);j++ {
		if loc, ok := locations[s[j]]; ok {
			i = max(i, loc+1)
		}
		maxLen = max(maxLen, j - i + 1)
		locations[s[j]] = j
	}
	return maxLen
}
