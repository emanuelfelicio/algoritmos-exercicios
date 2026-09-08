package algoritmosexercicios

func climbStairs(n int) int {
	var f func(int) int

	cache := make(map[int]int, 0)

	f = func(i int) int {
		if i <= 0 {
			if i == 0 {

				return 1
			}
			return 0
		}
		// verifica se valor esta em cache(memoization)
		if _, exist := cache[i]; exist {
			return cache[i]
		}
		
		cache[i] = f(i-1) + f(i-2)
		return cache[i]
	}

	return f(n)
}
