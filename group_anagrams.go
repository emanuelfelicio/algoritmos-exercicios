package algoritmosexercicios

func GroupAnagrams(strs []string) [][]string {
	// criar hashmap map[[26]rune][]string
	// itera sobre strs
	// aloca array [26]rune (chave da hashmap)
	// cada iteração conta letras(runes), exemplo: [1,2,0,0...]
	// adiciona string ao map, se slice existir inserir no final, se nao , aloca novo slice e adiciona
	m := make(map[[26]rune][]string)
	for i, word := range strs {
		var a [26]rune
		for _, r := range strs[i] {
			a[r-'a']++ // -'a' ajusta index
		}
		m[a] = append(m[a], word)
	}

	result := make([][]string, 0)

	for _, v := range m {
		result = append(result, v)
	}

	return result
}
