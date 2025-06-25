package arr

func mergeAlternately(word1 string, word2 string) string {
	minlen := len(word1)
	if minlen > len(word2) {
		minlen = len(word2)
	}
	left1 := word1[:minlen]
	left2 := word2[:minlen]
	left1Sub := word1[minlen:]
	left2Sub := word2[minlen:]
	font := ""
	for i, v := range left1 {
		a := string(v)
		b := string(left2[i])
		font += a + b
	}
	return font + left1Sub + left2Sub
}
