package algoritmosexercicios

import (
	"slices"
)

func CarFleet(target int, position []int, speed []int) int {
	s := make([][2]int, 0)

	for i := range position {
		s = append(s, [2]int{position[i], speed[i]})
	}

	slices.SortFunc(s, func(a, b [2]int) int {
		return b[0] - a[0]
	})
	results := make([]float64, 0)
	time := float64((target - s[0][0])) / float64(s[0][1])
	results = append(results, time)

	for i := 1; i < len(s); i++ {
		time = float64((target - s[i][0])) / float64(s[i][1])

		if time > results[len(results)-1] {
			results = append(results, time)
		}
	}

	return len(results)

}
