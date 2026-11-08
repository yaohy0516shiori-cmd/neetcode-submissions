func isValidSudoku(board [][]byte) bool {
	// noDup reports whether 9 cells contain no repeated digit. '.' is ignored.
	noDup := func(cells [9]byte) bool {
		seen := [9]bool{}
		for _, ch := range cells {
			if ch == '.' {
				continue
			}
			d := ch - '1' // '1'..'9' -> 0..8
			if seen[d] {
				return false
			}
			seen[d] = true
		}
		return true
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			// first time we reach row i: check the whole row
			if j == 0 {
				var row [9]byte
				for k := 0; k < 9; k++ {
					row[k] = board[i][k]
				}
				if !noDup(row) {
					return false
				}
			}

			// first time we reach column j: check the whole column
			if i == 0 {
				var col [9]byte
				for k := 0; k < 9; k++ {
					col[k] = board[k][j]
				}
				if !noDup(col) {
					return false
				}
			}

			// top-left corner of a 3x3 box: check the whole box
			if i%3 == 0 && j%3 == 0 {
				var box [9]byte
				t := 0
				for m := i; m < i+3; m++ {
					for n := j; n < j+3; n++ {
						box[t] = board[m][n]
						t++
					}
				}
				if !noDup(box) {
					return false
				}
			}
		}
	}
	return true
}