func isAnagram(s string, t string) bool {
	letters_counts_first := countLettersInWord(s)
	letters_counts_second := countLettersInWord(t)

	for letter, count := range letters_counts_first {
		if (letters_counts_second[letter] != count) {
			return false
		}
		delete(letters_counts_second, letter)
	}
	if len(letters_counts_second) != 0 {
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
