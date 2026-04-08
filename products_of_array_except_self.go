package algoritmosexercicios

func ProductExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	// preenche com 1 , elemento neutro da multiplicacao
	for i := range res {
		res[i] = 1
	}

	prefix := 1
	// exemplo: nums[4,3,5,6]
	// res[i] = numero da esquerda multiplicados, sem incluir res[i]
	for i := 0; i < len(nums); i++ {
		res[i] = prefix
		prefix *= nums[i]
	}

	postfix := 1
	//nums[4,3,5,6]
	//res[i] = res[i] * numeros da direta multiplicados
	//res[90,120,72,60]
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= postfix
		postfix *= nums[i]
	}

	return res
}
