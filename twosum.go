package algoritmosexercicios

func twoSum(nums []int, target int) []int {

    s := make(map[int]int)

    for currentIndex := range nums {

        // calcula par complementar
        complement := target - nums[currentIndex]

        // verifica se existe no hashmap
        if complementIndex, ok := s[complement]; ok {
            return []int{complementIndex, currentIndex}
        }

        s[nums[currentIndex]] = currentIndex
    }

    return []int{}
}