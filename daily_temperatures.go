package algoritmosexercicios

func dailyTemperatures(temperatures []int) []int {
	// array preenchido com zeros
	res := make([]int, len(temperatures))
	stack := []int{}

	for index, curTemp := range temperatures {
		top := temperatures[stack[len(stack)-1]]
		for len(stack) > 0 && curTemp > top {
			//pop
			stackInd := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			
			res[stackInd] = index - stackInd
		}
		//push indice na stack
		stack = append(stack, index)
	}

	return res
}
