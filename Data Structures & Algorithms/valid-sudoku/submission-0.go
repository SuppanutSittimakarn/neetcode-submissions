
func isValidSudoku(board [][]byte) bool {
	// check rows
	for i := range 9 {
		mp := make(map[byte]int)
		for k := range 9 {
			s := board[i][k]
			if _, f := mp[s]; f {
				return false
			} else if s != 46 {
				mp[s]++
			}
		}
	}

	// check cols
	for i := range 9 {
		mp := make(map[byte]int)
		for k := range 9 {
			s := board[k][i]
			if _, f := mp[s]; f {
				return false
			} else if s != 46 {
				mp[s]++
			}
		}
	}

	for square := range 9 {
		mp := make(map[byte]int)
		for i := range 3 {
			for j := range 3 {
				row := (square/3)*3 + i
				col := (square%3)*3 + j
				s := board[row][col]
				if _, f := mp[s]; f {
					return false
				} else if s != 46 {
					mp[s]++
				}
			}
		}
	}

	return true
}
