func isValidSudoku(board [][]byte) bool {
	var rows, cols, boxes [9][9]bool
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			ch := board[i][j]
			if ch == '.' {
				continue
			}
			d := ch - '1'
			b := (i/3)*3 + j/3
			if rows[i][d] || cols[j][d] || boxes[b][d] {
				return false
			}
			rows[i][d], cols[j][d], boxes[b][d] = true, true, true
		}
	}
	return true
}
