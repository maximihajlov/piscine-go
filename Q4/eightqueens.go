package Q4

import (
	"fmt"

	"github.com/01-edu/z01"
)

var board [8][8]int

func EightQueens() {
	setQueen(0)
}

func setQueen(x int) {
	if x == 8 {
		showBoard()
		showSolution()
		z01.PrintRune('\n')
	}
	for i := 0; i < 8; i++ {
		if tryQueen(x, i) {
			board[x][i] = 1
			setQueen(x + 1)
			board[x][i] = 0
		}
	}
}

func tryQueen(x, y int) bool {
	for i := 0; i < x; i++ {
		if board[i][y] != 0 {
			return false
		}
	}
	for i := 0; i <= x && y-i >= 0; i++ {
		if board[x-i][y-i] != 0 {
			return false
		}
	}
	for i := 1; i <= x && y+i < 8; i++ {
		if board[x-i][y+i] != 0 {
			return false
		}
	}
	return true
}

func showSolution() {
	for i := 0; i < 8; i++ {
		for j := '0'; j < '8'; j++ {
			if board[i][j-'0'] == 1 {
				z01.PrintRune(j + 1)
			}
		}
	}
}

func showBoard() {
	for i := 0; i < 8; i++ {
		fmt.Print(board[i])
		fmt.Println()
	}
}
