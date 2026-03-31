package algoritmosexercicios

type NumArray struct {
	nums     []int
	sumArray []int
}

func Constructor(nums []int) NumArray {
	s := make([]int, len(nums))

	if len(nums) <= 0 {
		return NumArray{}
	}

	s[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		s[i] = s[i-1] + nums[i]
	}

	return NumArray{nums: nums, sumArray: s}
}

func (this *NumArray) SumRange(left int, right int) int {
	// aproach 1 calcula a soma do range por chamada
	// sum := 0
	// for i := left; i <= right; i++ {
	// 	sum += this.nums[i]
	// }
	// return sum

	// aproach 2 calcula com base no array de acúmulo pré-calculado no construtctor
	if left == 0 {
		return this.sumArray[right]
	}

	return this.sumArray[right] - this.sumArray[left-1]

}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
