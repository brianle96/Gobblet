package objects

import (
	"fmt"
	"strconv"
)

//
//
type Board struct {
	Grid [3][3] Piece
	Turn string
}
//
//
func (b Board) GetCurrentBoard() [3][3] Piece {
	return b.Grid
}
//
//[row][col]
func (b Board) UpdateBoard(move Move) Board {
	piece := Piece{Color: move.Color, Size: move.Size}
	b.Grid[move.Row][move.Col] = piece
	if move.Color == "brown" {
		b.Turn = "green"
	} else {
		b.Turn = "brown"
	}
	return b
}
//
//
func (b Board) PrintCurrentBoard() {
	fmt.Println()
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			currColor := b.Grid[i][j].Color
			currSize := strconv.Itoa(b.Grid[i][j].Size)
			if j != 2 {
				fmt.Printf("%v,%v|",currColor, currSize)
			} else {
				fmt.Printf("%v,%v",currColor, currSize)
			}
		}
		//
		if i != 2 {
			fmt.Println("\n-----------------------")
		} else {
			fmt.Println()
		}
	}
}
//
//
func (b Board) CheckForWin() bool {
	winCondition := false
	//
	if b.CheckLeftDiagonal() {
		winCondition = true
	} else if b.CheckRightDiagonal() {
		winCondition = true
	}
	//
	for i :=0 ; i<= 2; i++ {
		if b.CheckHorizontal(i) {
			winCondition = true
		} else if b.CheckVertical(i) {
			winCondition = true
		}
	}
	return winCondition
}
//
//
func (b Board) CheckHorizontal(i int) bool {
	winCondition := false
	first := b.Grid[i][0].Color
	second := b.Grid[i][1].Color
	third := b.Grid[i][2].Color
	if first == second && second == third && first == third && first != "empty" {
		winCondition = true
	}
	return winCondition
}
//
//
func (b Board) CheckVertical(i int) bool {
	winCondition := false
	first := b.Grid[0][i].Color
	second := b.Grid[1][i].Color
	third := b.Grid[2][i].Color
	if first == second && second == third && first == third && first != "empty" {
		winCondition = true
	}
	return winCondition
}
//
//
func (b Board) CheckLeftDiagonal() bool {
	winCondition := false
	first := b.Grid[0][0].Color
	second := b.Grid[1][1].Color
	third := b.Grid[2][2].Color
	if first == second && second == third && first == third && first != "empty" {
		winCondition = true
	}
	return winCondition
}
//
//
func (b Board) CheckRightDiagonal() bool {
	winCondition := false
	first := b.Grid[0][2].Color
	second := b.Grid[1][1].Color
	third := b.Grid[2][0].Color
	if first == second && second == third && first == third && first != "empty" {
		winCondition = true
	}
	return winCondition
}
//
//
func InitBoard(turn string) Board {
	gameBoard := Board{Turn: turn}
	blankPiece := Piece{Color: "empty", Size: 0}
	tempBoard := [3][3] Piece {{blankPiece,blankPiece,blankPiece},
							   {blankPiece,blankPiece,blankPiece},
							   {blankPiece,blankPiece,blankPiece}}
	gameBoard.Grid = tempBoard
	return gameBoard
}