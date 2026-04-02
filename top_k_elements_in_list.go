package algoritmosexercicios

import (
	"slices"
)

func topKFrequent(nums []int, k int) []int {
	return solution1(nums, k)
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
