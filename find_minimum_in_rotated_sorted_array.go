package algoritmosexercicios

// rotacionar len(nums) vezes mantém ordem original
// elementos estavam ordenados antes da rotação
// 1 2 3 4 -> 4,1,2,3 - > 3,4,1,2 -> 2,3,4,1
func findMin(nums []int) int {
	l := 0
	r := len(nums) - 1

	for l < r {
		m := (l + r) / 2
		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}

	return nums[l]
}
