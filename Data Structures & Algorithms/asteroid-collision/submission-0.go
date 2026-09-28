func asteroidCollision(asteroids []int) []int {
	res:=[]int{}

	for i :=range asteroids {

		asteroid:=asteroids[i]

		for len(res)>0 && asteroid <0 && res[len(res)-1]>0{
			collision:=asteroid+res[len(res)-1]

			if collision==0 {
				asteroid=0
				res=res[:len(res)-1]
			} else if collision >0 {
				asteroid=0
			} else {
				res=res[:len(res)-1]
			}

		}

		if asteroid!=0 {
			res=append(res,asteroid)
		}
	}

	return res
}
