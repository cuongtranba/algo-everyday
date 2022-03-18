package twopointer

func mergeAlternately(word1 string, word2 string) string {
	res := ""
	loop := len(word1)
	if len(word2) > loop {
		loop = len(word2)
	}

	for i := 0; i < loop; i++ {
		if i > len(word1)-1 {
			for j := i; j < len(word2); j++ {
				res += string(word2[j])
			}
			break
		} else if i > len(word2)-1 {
			for j := i; j < len(word1); j++ {
				res += string(word1[j])
			}
			break
		} else {
			res += string(word1[i]) + string(word2[i])
		}
	}

	return res
}
