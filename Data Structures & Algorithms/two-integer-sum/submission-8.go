func twoSum(nums []int, target int) []int {
	ids := make(map[int]int)
	for i, n := range nums {
		if j, ok := ids[target - n]; ok {
			return []int{j, i}
		}
		ids[n] = i
	}
	return []int{0, 0}
}
