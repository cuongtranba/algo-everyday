package prefixsum

func maxScore(s string) int {
	total1 := 0
	result := 0
	for _, v := range s {
		if v == '1' {
			total1 += 1
		}
	}
	total0 := 0
	for _, v := range s {
		if v == '1' {
			total1--
		} else {
			total0++
		}
		score := total0 + total1
		result = max(result, score)
	}
	if total0 == 0 || total1 == 0 {
		return 1
	}
	return result
}
