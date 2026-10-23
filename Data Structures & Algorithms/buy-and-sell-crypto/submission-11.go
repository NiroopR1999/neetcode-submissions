func maxProfit(prices []int) int {
	buy,sell,profit:=0,0,0

	for sell<len(prices) {
		if prices[sell]<prices[buy] {
			buy=sell
		} else {
			profit=max(profit,prices[sell]-prices[buy])
			sell++
		}
	}
	return profit
}
