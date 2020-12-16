package dynamic

//买卖股票的最佳时机
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	min := prices[0]
	max := 0
	for i := 0; i < len(prices); i++ {
		if min > prices[i] {
			min = prices[i]
			continue
		}
		if prices[i]-min > max {
			max = prices[i] - min
		}
	}
	return max
}
