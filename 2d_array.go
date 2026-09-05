package algoritmosexercicios

import "math"

// array 6x6
// calcula bloco a cada colun:
// X X X
//   X
// X X X	
func hourglassSum(arr [][]int32) int32 {
	var maxSum int32 = math.MinInt32

	for i := range 4 {
		for j := range 4 {
			sum := arr[i][j] + arr[i][j+1] + arr[i][j+2] + arr[i+1][j+1] + arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
			maxSum = max(sum, maxSum)
		}

	}
	return maxSum

}
