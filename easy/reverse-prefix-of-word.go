func reversePrefix(word string, ch byte) string {
	found := false
	wordRev := ""
	for _, w := range word {
		if found == true {
			wordRev = wordRev + string(w)
		} else {
			wordRev = string(w) + wordRev
		}

		if w == rune(ch) {
			found = true
		}
	}

	if found == true {
		return wordRev
	}

	return word
}