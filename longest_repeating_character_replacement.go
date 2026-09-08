package algoritmosexercicios

func characterReplacement(s string, k int) int {
	res := 0
	charSet := make(map[byte]struct{})

	for i := 0; i < len(s); i++ {
		charSet[s[i]] = struct{}{}
	}

	// itera pro cada letra distinta em s
	for c := range charSet {
		maxFreq, l := 0, 0
		for r := 0; r < len(s); r++ {
			if s[r] == c {
				maxFreq++
			}

			// move left pointer até a condição ser válida
			for (r-l+1)-maxFreq > k {
				if s[l] == c {
					maxFreq--
				}
				l++
			}

			res = max(res, r-l+1)
		}
	}

	return res
}
