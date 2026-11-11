func searchMatrix(matrix [][]int, target int) bool {
	rows:=len(matrix)
	cols:=len(matrix[0])
	left,right:=0,(rows*cols)-1

	for left<=right{
		mid:=(left+right)/2
		row:=mid/cols
		col:=mid%cols
		if target == matrix[row][col] {
			return true
		} else if target > matrix[row][col] {
			left=mid+1
		} else {
			right=mid-1
		}
	}

	return false

}
