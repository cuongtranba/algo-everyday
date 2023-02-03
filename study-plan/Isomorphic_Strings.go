package studyplan

func isIsomorphic(s string, t string) bool {
	keyPairS := make(map[string]string)
	keyPairT := make(map[string]string)

	for i := 0; i < len(s); i++ {
		if _, ok := keyPairT[string(t[i])]; !ok {
			keyPairT[string(t[i])] = string(s[i])
		} else {
			if keyPairT[string(t[i])] != string(s[i]) {
				return false
			}
		}

		if _, ok := keyPairS[string(s[i])]; !ok {
			keyPairS[string(s[i])] = string(t[i])
		} else {
			if keyPairS[string(s[i])] != string(t[i]) {
				return false
			}
		}
	}
	return true
}
