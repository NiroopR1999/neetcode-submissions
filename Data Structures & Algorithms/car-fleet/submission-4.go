func carFleet(target int, position []int, speed []int) int {
	cars:=[][2]int{}
	n:=len(position)
	for i:=range n {
		cars=append(cars,[2]int{position[i],speed[i]})
	}

	sort.Slice(cars, func (i, j int) bool {
		return cars[i][0]>cars[j][0]
	})

	fleet:=[]float64{}

	for _,car :=range cars {
		time:=float64(target-car[0]) / float64(car[1])

		if len(fleet)==0 || time >fleet[len(fleet)-1] {
			fleet=append(fleet,time)
		}
	}

	return len(fleet)
}
