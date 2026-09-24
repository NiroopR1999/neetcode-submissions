func maxProfit(prices []int) int {
	if len(prices)==0 {
		return 0
	}
	maxProfit,buy,sell:=0,0,1

	for sell<len(prices) {
		if prices[sell]<prices[buy] {
			buy=sell
			sell++
			continue
		}
		profit:=prices[sell]-prices[buy]
		if profit>maxProfit{
			maxProfit=profit
		}
		sell++
	}

	return maxProfit
}
