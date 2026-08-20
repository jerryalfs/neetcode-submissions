func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minPrice := prices[0]
	maxProfit := 0
	for i:=1; i < len(prices); i++ {
		priceToday := prices[i]
		currentProfit := priceToday - minPrice
		if currentProfit > maxProfit {
			maxProfit = currentProfit
		}
		minPrice = min(minPrice,priceToday)
	}
	return maxProfit
}