package main

import (
	"errors"
	"fmt"
	"log"
)

type Board struct {
	size      int
	positions [][]string
}

func NewBoard(size int) *Board {
	positions, err := initBoard(size)
	if err != nil {
		log.Fatalln(err)
	}

	return &Board{
		positions: positions,
		size:      size,
	}
}

func initBoard(size int) ([][]string, error) {
	if size <= 0 {
		return nil, errors.New("Size must be positive")
	}

	pos := make([][]string, size)
	for i := 0; i < size; i++ {
		pos[i] = make([]string, size)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			pos[i][j] = "."
		}
	}
	return pos, nil
}

func (b *Board) Size() int {
	return b.size
}

func (b *Board) Place(r int, c int) {
	if r < 0 || r >= b.size || c < 0 || c >= b.size {
		log.Fatalln("indexes must be in-bound")
	}
	b.positions[r][c] = "Q"
}

func (b *Board) Remove(r int, c int) {
	if r < 0 || r >= b.size || c < 0 || c >= b.size {
		log.Fatalln("indexes must be in-bound")
	}
	b.positions[r][c] = "."
}

func (b *Board) Print() {
	for i := 0; i < b.size; i++ {
		fmt.Printf("%v\n", b.positions[i])
	}
	fmt.Println("")
}

func (b *Board) IsSafe(r int, c int) bool {
	if r < 0 || r >= b.size || c < 0 || c >= b.size {
		log.Fatalln("indexes must be in-bound")
	}

	// check row left side of r
	for i := 0; i <= c; i++ {
		if b.positions[r][i] == "Q" {
			return false
		}
	}

	// check upper diagonal on left side of r
	for i, j := r, c; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if b.positions[i][j] == "Q" {
			return false
		}
	}

	// check lower diagonal on left side of r
	for i, j := r, c; i < b.size && j >= 0; i, j = i+1, j-1 {
		if b.positions[i][j] == "Q" {
			return false
		}
	}

	return true
}
