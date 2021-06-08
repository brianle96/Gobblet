package modules
//
import (
	"fmt"
	"strconv"
)
//
type Board struct {
	Pieces1 [] Piece
	Pieces2 [] Piece
	Grid [3][3] Piece
	Turn bool
}
//
//
func (b Board) PrintCurrentBoard() {
	fmt.Println()
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			currPlayer := b.Grid[i][j].Player
			currSize := strconv.Itoa(b.Grid[i][j].Size)
			if j != 2 {
				fmt.Printf("%v,%v|",currPlayer, currSize)
			} else {
				fmt.Printf("%v,%v",currPlayer, currSize)
			}
		}
		//
		if i != 2 {
			fmt.Println("\n--------------------")
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
	first := b.Grid[i][0].Player
	second := b.Grid[i][1].Player
	third := b.Grid[i][2].Player
	if first == second && second == third && first == third && first != "none" {
		winCondition = true
	}
	return winCondition
}
//
//
func (b Board) CheckVertical(i int) bool {
	winCondition := false
	first := b.Grid[0][i].Player
	second := b.Grid[1][i].Player
	third := b.Grid[2][i].Player
	if first == second && second == third && first == third && first != "none" {
		winCondition = true
	}
	return winCondition
}
//
//
func (b Board) CheckLeftDiagonal() bool {
	winCondition := false
	first := b.Grid[0][0].Player
	second := b.Grid[1][1].Player
	third := b.Grid[2][2].Player
	if first == second && second == third && first == third && first != "none" {
		winCondition = true
	}
	return winCondition
}
//
//
func (b Board) CheckRightDiagonal() bool {
	winCondition := false
	first := b.Grid[0][2].Player
	second := b.Grid[1][1].Player
	third := b.Grid[2][0].Player
	if first == second && second == third && first == third && first != "none" {
		winCondition = true
	}
	return winCondition
}
//
//
//[row][col]
func (b Board) UpdateBoard(move Move) Board {
	piece := Piece{Player: move.Player, Size: move.Size}
	b.Grid[move.Row][move.Col] = piece
	if move.Player == PLAYER1_COLOR && !b.IsUnplayable(b.Pieces2) {
		b.Turn = false
	} else if move.Player == PLAYER2_COLOR && !b.IsUnplayable(b.Pieces1){
		b.Turn = true
	}
	return b
}
//
//
func (b Board) EmptyTeams() bool {
	if b.Pieces1 == nil && b.Pieces2 == nil {
		return true	
	}
	return false
}
//
//
func (b Board) IsUnplayable(currPieces []Piece) bool {
	//
	unplayableSpaces := 0
	isBlocked := false
	for _, piece := range currPieces {
		for i:=0; i <= 2; i++ {
			for j:=0; j <= 2; j++ {
				currBoardSize := b.Grid[i][j].Size
				if piece.Size <= currBoardSize {
					unplayableSpaces += 1
				}
			}
		}
	}
	failedNumber := len(currPieces)*9
	if unplayableSpaces == failedNumber {
		isBlocked = true
	} else {
		isBlocked = false
	}
	return isBlocked
}