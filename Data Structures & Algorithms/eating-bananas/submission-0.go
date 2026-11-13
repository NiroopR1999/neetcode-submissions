func minEatingSpeed(piles []int, h int) int {
	maxBanana:=0

	for _, pile:=range piles {
		if pile>maxBanana{
			maxBanana=pile
		}
	}

	left,right:=1,maxBanana

	for left<right {
		mid:=(left+right)/2
		hours:=0
		for _,pile:=range  piles {
			hours+=(pile+mid-1)/mid
		}
		if hours<=h {
			right=mid
		} else {
			left=mid+1
		}
	}
	return right
}
