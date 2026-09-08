func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
        return false
    }

	letters_counts_first := countLettersInWord(s)

	for _, letter := range t {
		letters_counts_first[letter]--
		if (letters_counts_first[letter] < 0) {
			return false
		}
		if (letters_counts_first[letter] == 0) {
			delete(letters_counts_first, letter)
		} 
	}
	if (len(letters_counts_first) != 0) {
		return false
	}
	return true
}

func countLettersInWord(word string) map[rune]int {
	letters_counts := make(map[rune]int)
	for _, letter := range word {
		letters_counts[letter]++
	}
	return letters_counts
}
