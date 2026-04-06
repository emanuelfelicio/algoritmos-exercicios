package algoritmosexercicios

import (
	"slices"
)

func TopKFrequent(nums []int, k int) []int {

	//segunda solucao : bucket sort
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	// aloca bucket  onde cada indice s[value] armazena elementos que
	// apareceram 'value' vezes
	s := make([][]int, len(nums)+1)
	for key, value := range m {
		s[value] = append(s[value], key)
	}

	res := make([]int, 0)
	// itera de trás para frente nos buckets e monta res
	for i := len(s) - 1; i >= 0; i-- {
		for _, v := range s[i] {
			res = append(res, v)
			if len(res) == k {
				return res
			}
		}
	}

	return res
}

// primeira solução: contar repetição . tempo O(n log n)
func solution1(nums []int, k int) []int {
	// conta frequencia de cada elemento
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	// armazena key,valu em array e ordena com base no valor
	a := make([][2]int, 0, len(m))
	for num, count := range m {
		a = append(a, [2]int{num, count})
	}
	slices.SortFunc(a, func(a, b [2]int) int {
		return a[1] - b[1]
	})

	// monta resposta
	res := make([]int, 0, k)

	for i := len(a) - 1; i >= len(a)-k; i-- {
		res = append(res, a[i][0])
	}

	return res
}
