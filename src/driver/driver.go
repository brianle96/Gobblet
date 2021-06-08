package main

//
//
import (
	"fmt"
	"strings"
	m "github.com/modules"
)

//
//
func main() {
	//
	noMoreMoves := false
	//
	winCondition := false
	//
	GameBoard := m.InitBoard()
	//
	for winCondition == false {
		//
		GameBoard.PrintCurrentBoard()
		//
		currPieces := m.GetCurrentPieces(GameBoard)
		//
		currMove := m.GetMove(currPieces)
		//
		legalMove, err := m.IsLegalMove(GameBoard, currMove, currPieces)
		//
		if legalMove {
			//
			GameBoard = m.RemovePiece(GameBoard, currMove)
			//
			GameBoard = GameBoard.UpdateBoard(currMove)
		} else {
			fmt.Printf("***********\nInvalid move!\n%v\n***********\n", err)
		}
		//
		winCondition = GameBoard.CheckForWin()
		//
		if GameBoard.EmptyTeams() {
			break
		} else if (GameBoard.IsUnplayable(GameBoard.Pieces1) && 
		GameBoard.IsUnplayable(GameBoard.Pieces2)) {
			noMoreMoves = true
			break
		}
	}
	//
	GameBoard.PrintCurrentBoard()
	//
	if GameBoard.EmptyTeams() {
		fmt.Println("\nRAN OUT OF PIECES! TIE! TRY AGAIN NEXT TIME")
	} else if noMoreMoves {
		fmt.Println("\nNO MORE VALID MOVES! TIE! TRY AGAIN NEXT TIME")
	} else if GameBoard.Turn {
		// since the "current" turn is on true, it would
		// mean that the last move made was false
		// therefore they won the match
		fmt.Printf("\n%v WON THE MATCH\n", strings.ToUpper(m.PLAYER2_COLOR))
	} else {
		fmt.Printf("\n%v WON THE MATCH\n", strings.ToUpper(m.PLAYER1_COLOR))
	}
}
