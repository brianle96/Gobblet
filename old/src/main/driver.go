package main
//
import (
	"fmt"
	o "github.com/objects"
)
//
func main() {
	//
	winCondition := false
	//
	Player1, Player2 := o.GeneratePair("Brian","Erin")
	//
	//
	GameBoard := o.InitBoard("brown")
	//
	for winCondition == false {
		// on top of currMove, do currPlayer
		currPlayer := o.Player{}
		currMove := o.Move{}
		//
		//
		GameBoard.PrintCurrentBoard()
		fmt.Printf("\n%v's Move!",GameBoard.Turn)
		if GameBoard.Turn == "brown" {
			currPlayer = Player1
		} else {
			currPlayer = Player2
		}
		fmt.Printf("\n%v's Team: %v", currPlayer.Color, currPlayer.Team)
		//
		//
		currMove = currPlayer.GetMoveFromHuman()
		//
		if o.IsLegalMove(currMove, GameBoard, currPlayer) {
			//
			currPlayer.Team = currPlayer.RemovePiece(currMove)
			// need to update the player objects
			if GameBoard.Turn == "brown" {
				Player1 = currPlayer
			} else {
				Player2 = currPlayer
			}
			// update the board with the move
			GameBoard = GameBoard.UpdateBoard(currMove)
		} else {
			fmt.Printf("***********\nInvalid move!\n***********\n")
		}
		//
		//
		winCondition = GameBoard.CheckForWin()
		if Player1.Team == nil && Player2.Team == nil {
			break
		}
	}
	//
	GameBoard.PrintCurrentBoard()
	//
	if Player1.Team == nil && Player2.Team == nil {
		fmt.Printf("\nTIE\n")
	} else if GameBoard.Turn == "brown" {
		fmt.Printf("\nGreen Won!\n")
	} else if GameBoard.Turn == "green"{
		fmt.Printf("\nBrown Won!\n")
	}
}