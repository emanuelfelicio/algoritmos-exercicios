package algoritmosexercicios

import "slices"

func threeSum(nums []int) [][]int {

	slices.Sort(nums)
	result := make([][]int, 0)

	for i := 0; i < len(nums)-2; i++ {
		// evita duplicatas na primeia posicao
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1

		for left < right {
			sum := nums[left] + nums[right] + nums[i]
			if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			} else {
				result = append(result, []int{nums[left], nums[right], nums[i]})

				// evita duplicatas na segunda e terceira posicao
				left++
				for nums[left] == nums[left-1] && left < right {
					left++
				}
			}
		}
	}
	return result
}
