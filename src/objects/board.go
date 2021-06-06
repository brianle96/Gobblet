package objects

import (
	"fmt"

	"strconv"
)

//
//
type Board struct {
	Grid [3][3] Piece
}
//
//
func (b Board) PrintCurrentBoard() {
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
			fmt.Println("\n-----------")
		} else {
			fmt.Println()
		}
	}
}
//
//
func InitBoard() Board {
	gameBoard := Board{}
	blankPiece := Piece{Color: "e", Size: 0}
	tempBoard := [3][3] Piece {{blankPiece,blankPiece,blankPiece},
							   {blankPiece,blankPiece,blankPiece},
							   {blankPiece,blankPiece,blankPiece}}
	gameBoard.Grid = tempBoard
	return gameBoard
}