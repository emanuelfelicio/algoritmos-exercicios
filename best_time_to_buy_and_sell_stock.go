package algoritmosexercicios

func maxProfit(prices []int) int {
	// [10,1,10,1,0,1] escolher menor e depois o maior depois dele, não funciona, nesse exemplo 
	// [10,1,10,1,0,1] proximo < atual; atual = proximo. senao calcular lucro

	profit := 0
	buy := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < buy {
			buy = prices[i]
		} else {
			profit = max(profit, prices[i]-buy)
		}
	}

	return profit
}
