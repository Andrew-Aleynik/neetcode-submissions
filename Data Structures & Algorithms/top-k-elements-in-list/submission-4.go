func topKFrequent(nums []int, k int) []int {
	frequencies := make(map[int]int)
	for _, num := range nums {
		frequencies[num]++
	}
	frequency_buckets := make([][]int, len(nums) + 1)
	for num, freq := range frequencies {
		frequency_buckets[freq] = append(frequency_buckets[freq], num)
	}
	answer := make([]int, k)
	i := 0
	for freq := len(frequency_buckets) - 1; freq >= 0; freq-- {
		for _, num := range frequency_buckets[freq] {
			answer[i] = num
			i++
			if i == k {
				return answer
			}
		}
	}
	return answer
}