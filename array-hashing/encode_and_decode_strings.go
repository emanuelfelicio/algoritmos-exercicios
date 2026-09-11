package algoritmosexercicios

import (
	"strconv"
)

type Solution struct {
}

func (s *Solution) Encode(strs []string) string {
	var result []byte

	for _, rawString := range strs {
		// adiciona tamanho e caractere de marcação
		// caso adicione somente o marcador ou tamanho, iria falhar em casos no qual caractere separador é válido.
		// padrão : 2&ab3&abc => [ab,abc]
		result = append(result, []byte(strconv.Itoa(len(rawString)))...)
		result = append(result, '&')

		result = append(result, []byte(rawString)...)
	}

	return string(result)
}

func (s *Solution) Decode(encoded string) []string {
	var result []string
	i := 0

	for i < len(encoded) {
		// ler tamanho até o marcador
		j := i
		for encoded[j] != '&' {
			j++
		}

		length, _ := strconv.Atoi(encoded[i:j])

		// extrai string
		start := j + 1 // pula marcador 
		end := start + length

		result = append(result, encoded[start:end])

		i = end 
	}

	return result
}
