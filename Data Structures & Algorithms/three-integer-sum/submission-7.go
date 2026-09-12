func threeSum(nums []int) [][]int {
	answer := make([][]int, 0)
	sort.Ints(nums)
	i := 0
	for ; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j, k := i+1, len(nums)-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				answer = append(answer, []int{nums[i], nums[j], nums[k]})
				for ;j < k && nums[j] == nums[j+1]; {
					j++
				}
				for ;j < k && nums[k] == nums[k-1]; {
					k--
				}
				j++
				k--
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
	}
	return answer
}
