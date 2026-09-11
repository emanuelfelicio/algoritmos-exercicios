package algoritmosexercicios

func twoSum2(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	res := make([]int, 0, 2)
	// ajusta 'ponteiros' conforme comparação com target
	// ate encontrar
	for left < right {
		if numbers[left]+numbers[right] > target {
			right--
		} else if numbers[left]+numbers[right] < target {
			left++
		} else {
			// ajusta para 1-indexed
			left, right = left+1, right+1
			res = append(res, left, right)
			break
		}
	}

	return res
}
