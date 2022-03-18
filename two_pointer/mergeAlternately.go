package twopointer

func mergeAlternately(word1 string, word2 string) string {
	res := ""
	loop := len(word1)
	if len(word2) > loop {
		loop = len(word2)
	}

	for i := 0; i < loop; i++ {
		if i > len(word1)-1 {
			res += word2[i:]
			break
		} else if i > len(word2)-1 {
			res += word1[i:]
			break
		} else {
			res += string(word1[i]) + string(word2[i])
		}
	}

	return res
}
