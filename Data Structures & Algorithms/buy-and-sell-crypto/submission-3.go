func maxProfit(prices []int) int {
	// check len price 0 or not, if zero just return 0
	lenP := len(prices)
	if lenP == 0 {
		return 0
	}
	// take first index as indetifier minimum price that we have
	// later on will get replace while looping if there's more lowest value
	minPrice := prices[0]
	maxProfit := 0
	// why we start i at 1 because want to compare min price that we have in hand with next index
	for i:=1; i < lenP; i++ {
		// assign prices current index to price today
		priceToday := prices[i]
		// profit today got from subtraction current day to previous day
		profitToday := prices[i] - minPrice
		// assign profit today to max profit if it's greater than
		if profitToday > maxProfit {
			maxProfit = profitToday
		}
		// assign min value as current price, in-case current index value was lowest than previous index
		minPrice = min(priceToday,minPrice)
	}
	return maxProfit
}