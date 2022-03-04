package slidingwindow

//Input: s = ""wobgrovw""
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	min := 0
	max := 1
	visit := make(map[string]string)
	total := 0
	for min < max {
		for i := min; i < max; i++ {
			if _, ok := visit[string(s[i])]; ok {
				min = min + 1
				if len(visit) > total {
					total = len(visit)
				}
				visit = make(map[string]string)
				break
			}
			visit[string(s[i])] = string(s[i])
			if max < len(s) {
				max = max + 1
			}
		}
	}
	if len(visit) > total {
		return len(visit)
	}
	return total
}
