package main

import (
	"fmt"
	"os"
	"strconv"
)

func NQueens(board *Board) {
	solveNQueens(board, 0)
}

func solveNQueens(board *Board, col int) {
	if col >= board.Size() {
		board.Print()
	} else {
		for row := 0; row < board.Size(); row++ {
			if board.IsSafe(row, col) {
				// chosoe
				board.Place(row, col)
				// explore
				solveNQueens(board, col+1)
				// unchoose
				board.Remove(row, col)
			}
		}
	}
}

func usage() {
	fmt.Println(`Usage: nqueens N
  N: number`)
	os.Exit(1)
}

func main() {
	if len(os.Args) <= 1 {
		usage()
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		usage()
	}

	b := NewBoard(n)
	NQueens(b)
}
