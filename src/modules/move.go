package modules
//
import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)
//
type Move struct {
	Size int
	Row int
	Col int
	Player string
}
//
func GetMove(currPieces []Piece) Move {
	currPlayer := currPieces[0].Player
	fmt.Printf("\n%v's Team: \n%v\n", currPlayer, currPieces)
	//
	fmt.Println("Please enter your move like so: (row, col, size):")
	var reader = bufio.NewReader(os.Stdin)
	message, _ := reader.ReadString('\n')
	//
	re := regexp.MustCompile("[0-9]+")
	//
	moveSlice := re.FindAllString(message, -1)
	//
	// _ means I don't want to store that value anywhere
	// this one is an err variable
	rowMove, _ := strconv.Atoi(moveSlice[0])
	colMove, _ := strconv.Atoi(moveSlice[1])
	sizeMove, _ := strconv.Atoi(moveSlice[2])
	//
	tempMove := Move{Size: sizeMove,
					Row: rowMove,
					Col: colMove,
					Player: currPlayer,
	}
	return tempMove
}
//
//
// should i allow a player to overlap their own piece
func IsLegalMove(gameBoard Board, move Move, currPieces []Piece) (bool, string) {
	if !isInBounds(move) {
		return false, "Not in Bounds"
	} else if !isRightSize(move, gameBoard) {
		return false, "Size Incorrect"
	} else if !isInTeam(move, currPieces) {
		return false, "Not in Team"
	}
	return true, ""
}
//
//
func isInBounds(move Move) bool{
	result := true
	if move.Row < 0 || move.Row >= 3 {
		result = false
	}
	if move.Col < 0 || move.Col >= 3 {
		result = false
	}
	return result
}
//
//
func isRightSize(move Move, gameBoard Board) bool {
	result := true
	intendedPosition := gameBoard.Grid[move.Row][move.Col].Size
	checkPlayer := gameBoard.Grid[move.Row][move.Col].Player
	if move.Size <= intendedPosition || checkPlayer == move.Player {
		result = false
	}
	return result
}
//
//
func isInTeam(move Move, currPieces []Piece) bool {
	result := false
	tempPiece := Piece{Player: move.Player,Size: move.Size}
	for _, copy := range currPieces{
		if copy == tempPiece {
			result = true
		}
	}
	return result
}