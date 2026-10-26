package algoritmosexercicios

func checkInclusion(s1 string, s2 string) bool {
	// O(26.N) => O(N)
	var solution1 = func(s1 string, s2 string) bool {
		if len(s1) > len(s2) {
			return false
		}

		a := [26]int{}

		for i := 0; i < len(s1); i++ {
			a[s1[i]-'a']++
		}

		m := make(map[[26]int]bool)
		m[a] = true

		chunk := len(s1)
		n := len(s2) - chunk + 1
		for i := range n {
			letters := [26]int{}

			for j := i; j < i+chunk; j++ {
				letters[s2[j]-'a']++
			}

			if ok := m[letters]; ok {
				return ok
			}
		}

		return false
	}

	_ = solution1(s1, s2)

	var solution2 = func(s1 string, s2 string) bool {
		if len(s1) > len(s2) {
			return false
		}

		letters := make(map[rune]int)

		for _, s := range s1 {
			letters[s]++
		}

		need := len(letters)

		chunk := len(s1)
		n := len(s2) - chunk + 1

		for i := range n {
			limit := i + chunk
			windowLetters := make(map[rune]int, chunk)
			cur := 0
			for _, r := range s2[i:limit] {

				windowLetters[r]++

				if letters[r] < windowLetters[r] {
					break
				}

				if letters[r] == windowLetters[r] {
					cur++
				}

				if need == cur {
					return true
				}
			}

		}

		return false
	}

	return solution2(s1, s2)

}
