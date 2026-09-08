func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]int][]string)
    for _, s := range strs {
        key := countLetters(s)
        groups[key] = append(groups[key], s)
    }
    result := make([][]string, 0, len(groups))
    for _, v := range groups {
        result = append(result, v)
    }
    return result
}

func countLetters(s string) [26]int{
	counts := [26]int{}
	for _, c := range s {
		counts[c - 'a']++
	}
	return counts
}
