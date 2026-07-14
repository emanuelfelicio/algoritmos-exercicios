package algoritmosexercicios

func largestRectangleArea(heights []int) int {
	maxArea := 0
	stack := make([][2]int, 0)

	for i, h := range heights {
		start := i
		// se o topo da stack for maior que a proxima altura 
		// pop e calcular area
		for len(stack) > 0 && stack[len(stack)-1][1] > h {
			//top() index and heigth
			index := stack[len(stack)-1][0]
			height := stack[len(stack)-1][1]
			//pop()
			stack = stack[:len(stack)-1]
			//calcula area
			area := height * (i - index)
			// computa maxArea
			if area > maxArea {
				maxArea = area
			}
			//extende o novo elemento ao index do elemento do topo da stack
			//por que height > h
			start = index
		}
		stack = append(stack, [2]int{start, h})
	}

	n := len(heights)
	// calcula area para o elementos que sobraram na stack
	for _, pair := range stack {
		area := pair[1] * (n - pair[0])
		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}
