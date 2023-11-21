package isogram

import "strings"

func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	counts := map[rune]int{}

	for _, ch := range word {
		if ch == ' ' || ch == '-' {
			continue
		}

		counts[ch] += 1
	}

	for _, count := range counts {
		if count > 1 {
			return false
		}
	}

	return true
}
