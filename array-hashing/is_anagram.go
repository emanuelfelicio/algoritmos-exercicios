package algoritmosexercicios

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	sMap := make(map[rune]int)
	tMap := make(map[rune]int)

	// preenche maps
	for i := 0; i < len(s); i++ {
		sMap[rune(s[i])]++
		tMap[rune(t[i])]++
	}

	if len(sMap) != len(tMap) {
		return false
	}
	// verifica se são iguais em chave e valor
	for k, v := range sMap {
		if tMap[k] != v {
			return false
		}
	}

	return true
}
