// Given a 0-indexed string word and a character ch, reverse the segment of word that starts at index 0 and ends at the index of the first occurrence of ch (inclusive). If the character ch does not exist in word, do nothing.

// For example, if word = "abcdefd" and ch = "d", then you should reverse the segment that starts at 0 and ends at 3 (inclusive). The resulting string will be "dcbaefd".

package twopointer

func reversePrefix(word string, ch byte) string {
	newStr := ""
	pos := -1
	for i := range word {
		if word[i] == ch {
			pos = i + 1
			for j := i; j >= 0; j-- {
				newStr += string(word[j])
			}
			break
		}
	}

	if pos == -1 {
		return word
	}

	for i := pos; i < len(word); i++ {
		newStr += string(word[i])
	}
	return newStr
}
