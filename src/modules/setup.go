package modules

//
import (
	"fmt"
	"math/rand"
	"time"
)

//
// true
var PLAYER1_COLOR string = "blue"
// false
var PLAYER2_COLOR string = "pink"
//
//
func InitBoard() Board {
	//
	//
	blankPiece := Piece{Player:"none", Size: 0}
	tempBoard := [3][3] Piece {{blankPiece,blankPiece,blankPiece},
							   {blankPiece,blankPiece,blankPiece},
							   {blankPiece,blankPiece,blankPiece}}
	//
	// rand uses a seed to randomize
	rand.Seed(time.Now().UnixNano())
	turns := []bool{true, false}
    randomIndex := rand.Intn(len(turns))
    whoStarts := turns[randomIndex]
	fmt.Println(randomIndex)
	//
	//
	gameBoard := Board{
		Pieces1: GenerateTeam(PLAYER1_COLOR),
		Pieces2: GenerateTeam(PLAYER2_COLOR),
		Grid: tempBoard,
		Turn: whoStarts,
	}
	return gameBoard
}
//
//
func GetCurrentPieces(gameBoard Board) [] Piece {
	playerTeam := []Piece{}
	if gameBoard.Turn {
		playerTeam = gameBoard.Pieces1
	} else {
		playerTeam = gameBoard.Pieces2
	}
	return playerTeam
}