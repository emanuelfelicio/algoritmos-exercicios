package algoritmosexercicios

func lengthOfLongestSubstring(s string) int {
	mp := make(map[byte]int)
	l, res := 0, 0

	for r := 0; r < len(s); r++ {
		if index, found := mp[s[r]]; found {
			// atualiza o inicio da nova janela para o proximo index válido
			l = max(index+1, l)
		}
		// 
		mp[s[r]] = r

		if r-l+1 > res {
			res = r - l + 1
		}
	}
	return res
}
