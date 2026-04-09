package algoritmosexercicios

func longestConsecutive(nums []int) int {

	set := make(map[int]struct{}, len(nums))
	// adiciona numero ao set
	for _, num := range nums {
		set[num] = struct{}{}
	}

	longest := 0
	for num := range set {
		// verifica 'num' é o primeiro da sequência
		if _, found := set[num-1]; !found {
			// conta todos os numeros existentes na sequencia
			length := 1
			for {
				if _, exists := set[num+length]; exists {
					length++
				} else {
					break
				}
			}
			// atualiza longest
			if length > longest {
				longest = length
			}
		}
	}
	return longest

}
