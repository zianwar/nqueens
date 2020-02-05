package main

func NQueens(board *Board) {
	NQueensHelper(board, 0)
}

func NQueensHelper(board *Board, col int) {
	if col >= board.Size() {
		board.Print()
	} else {
		for row := 0; row < board.Size(); row++ {
			if board.IsSafe(row, col) {
				// choosee
				board.Place(row, col)
				// explore
				NQueensHelper(board, col+1)
				// unchoose
				board.Remove(row, col)
			}
		}
	}
}

func main() {
	b := NewBoard(4)
	NQueens(b)
}
