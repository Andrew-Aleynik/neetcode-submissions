func maxProfit(prices []int) int {
	best_diff := 0
	for i := 0; i < len(prices)-1; i++ {
		for j := i+1; j < len(prices); j++ {
			diff := prices[j] - prices[i]
			if diff > best_diff {
				best_diff = diff
			} 
		}
	}
	return best_diff
}
