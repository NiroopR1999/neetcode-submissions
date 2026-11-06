func isValidSudoku(board [][]byte) bool {
	rows:=make([]map[byte]bool,9)
	cols:=make([]map[byte]bool,9)
	grids:=make([]map[byte]bool,9)

	for i:=0;i<9;i++ {
		rows[i]=make(map[byte]bool)
		cols[i]=make(map[byte]bool)
		grids[i]=make(map[byte]bool)
	}

	for row:=0;row<9;row++ {
		for col:=0;col<9;col++ {
			num:=board[row][col]
			if num=='.' {
				continue
			}
			gridIndex:=(row/3)*3+(col/3)
			if rows[row][num] || cols[col][num] || grids[gridIndex][num] {
				return false
			}
			rows[row][num]=true
			cols[col][num]=true
			grids[gridIndex][num]=true
		}
	}

	return true

}
