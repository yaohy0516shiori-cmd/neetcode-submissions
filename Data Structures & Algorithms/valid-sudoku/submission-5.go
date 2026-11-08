func isValidSudoku(board [][]byte) bool {
	noDup:=func(cell func(k int) byte)bool{
		seen:=[9]bool{}
		for k:=0;k<9;k++{
			ch:=cell(k)
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
	for i := 0; i < 9; i++ {
		i := i // 让闭包捕获当前值（Go 1.22 之前需要）
		if !noDup(func(k int) byte { return board[i][k] }) { // 行
			return false
		}
		if !noDup(func(k int) byte { return board[k][i] }) { // 列
			return false
		}
		bi, bj := (i/3)*3, (i%3)*3
		if !noDup(func(k int) byte { return board[bi+k/3][bj+k%3] }) { // 宫格
			return false
		}
	}
	return true
}

