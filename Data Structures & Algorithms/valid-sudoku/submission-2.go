func isValidSudoku(board [][]byte) bool {
	noDup := func(cells [9]byte) bool {
		seen := [9]bool{}
		for _, ch := range cells {
			if ch == '.' {
				continue
			}
			d := ch - '1'
			if seen[d] {
				return false
			}
			seen[d] = true
		}
		return true
	}
	for i := 0; i < 9; i++ { // 9 行
		var row [9]byte
		for k := 0; k < 9; k++ {
			row[k] = board[i][k]
		}
		if !noDup(row) {
			return false
		}
	}
	for j := 0; j < 9; j++ { // 9 列
		var col [9]byte
		for k := 0; k < 9; k++ {
			col[k] = board[k][j]
		}
		if !noDup(col) {
			return false
		}
	}
	for b := 0; b < 9; b++ { // 9 个宫格
		bi, bj := (b/3)*3, (b%3)*3 // 第 b 个宫格的左上角
		var box [9]byte
		for k := 0; k < 9; k++ {
			box[k] = board[bi+k/3][bj+k%3]
		}
		if !noDup(box) {
			return false
		}
	}
	return true
}
