package algoritmosexercicios

func isValidSudoku(board [][]byte) bool {
    rows := make([]map[byte]bool, 9)
    cols := make([]map[byte]bool, 9)
    squares := make([]map[byte]bool, 9)

	// aloca maps
    for i := range 9 {
        rows[i] = make(map[byte]bool)
        cols[i] = make(map[byte]bool)
        squares[i] = make(map[byte]bool)
    }

    for i := range 9 {
        for j := range 9 {
            if board[i][j] == '.' {
                continue
            }
            val := board[i][j]
            squareIdx := (i/3)*3 + j/3

            if rows[i][val] || cols[j][val] ||
               squares[squareIdx][val] {
                return false
            }

			//adiciona ao sets
            rows[i][val] = true
            cols[j][val] = true
            squares[squareIdx][val] = true
        }
    }

    return true
}