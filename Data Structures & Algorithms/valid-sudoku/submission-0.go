func isValidSudoku(board [][]byte) bool {
	check := func(i, j int) bool {
		// row i
		seen := [9]bool{}
		for k := 0; k < 9; k++ {
			ch := board[i][k]
			if ch == '.' {
				continue
			}
			d := ch - '1'
			if seen[d] {
				return false
			}
			seen[d] = true
		}
		// column j
		seen = [9]bool{}
		for k := 0; k < 9; k++ {
			ch := board[k][j]
			if ch == '.' {
				continue
			}
			d := ch - '1'
			if seen[d] {
				return false
			}
			seen[d] = true
		}
		// 3x3 box containing (i, j)
		seen = [9]bool{}
		bi, bj := i-i%3, j-j%3
		for m := bi; m < bi+3; m++ {
			for n := bj; n < bj+3; n++ {
				ch := board[m][n]
				if ch == '.' {
					continue
				}
				d := ch - '1'
				if seen[d] {
					return false
				}
				seen[d] = true
			}
		}
		return true
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if !check(i, j) {
				return false
			}
		}
	}
	return true
}
