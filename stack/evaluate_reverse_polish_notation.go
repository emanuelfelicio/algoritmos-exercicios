package algoritmosexercicios

import "strconv"

func evalRPN(tokens []string) int {

	// solução usando stack
	stack := make([]int, 0, len(tokens))

	for _, v := range tokens {
		switch v {
		case "+", "*", "-", "/":
			op1 := stack[len(stack)-1]
			op2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			var result int
			switch v {
			case "+":
				result = op1 + op2
			case "-":
				result = op2 - op1
			case "*":
				result = op1 * op2
			case "/":
				result = op2 / op1
			}
			stack = append(stack, result)

		default:
			v, _ := strconv.Atoi(v)
			stack = append(stack, v)
		}
	}
	return stack[0]

	// solução usando recursão 

	// var dfs func() int
	// index := len(tokens) - 1
	// dfs = func() int {
	// 	token := tokens[index]
	// 	index--

	// 	if token != "+" && token != "-" && token != "*" && token != "/" {
	// 		val, _ := strconv.Atoi(token)
	// 		return val
	// 	}

	// 	right := dfs()
	// 	left := dfs()

	// 	switch token {
	// 	case "+":
	// 		return left + right
	// 	case "-":
	// 		return left - right
	// 	case "*":
	// 		return left * right
	// 	default:
	// 		return left / right
	// 	}
	// }

	// return dfs()

}
