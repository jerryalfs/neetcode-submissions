func isValidSudoku(board [][]byte) bool {
    var rows,cols,boxes [9][9]bool
    for r :=0 ; r < 9; r ++ {
        for c :=0 ; c < 9; c ++ {
            // check if the value was '.' continue the loop
            if board[r][c] == '.' {
                continue
            }
            // d meaning it's digit
            // as identifier will subtract with anchor '1'
            // so we know the value of the cell
            d := board[r][c] - '1'
            // b meaning band, to determine the 
            // cell was located in which box
            // since it's separate per 3x3 
            // so going the divide r and c with 3
            b := (r/3)*3 + c/3
            // check the location already marked as seen or not
            if rows[r][d] || cols[c][d] || boxes[b][d] {
                return false
            }
            // mark the location with the digit already seen
            rows[r][d],cols[c][d],boxes[b][d] = true,true,true
        }
    }
    return true
}
