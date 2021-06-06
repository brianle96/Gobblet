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
func InitBoard(turn string) Board {
	gameBoard := Board{Turn: turn}
	blankPiece := Piece{Color: "empty", Size: 0}
	tempBoard := [3][3] Piece {{blankPiece,blankPiece,blankPiece},
							   {blankPiece,blankPiece,blankPiece},
							   {blankPiece,blankPiece,blankPiece}}
	gameBoard.Grid = tempBoard
	return gameBoard
}
//
//
func CheckForWin() {
	for i :=1 ; i<= 3; i++ {
		
	}
}
//
//
func CheckHorizontal() bool {
	return true
}
//
//
func CheckVertical() bool {
	return true
}
//
//
func CheckDiagonal() bool {
	return true
}