package main

import (
	"fmt"
	o "github.com/objects"
)

func main() {
	//
	winCondition := false
	//
	Player1, Player2 := o.GeneratePair("Brian","Erin")
	//
	// playerOneName := Player1.GetPlayerName()
	// playerTwoName := Player2.GetPlayerName()
	// fmt.Println(playerOneName)
	// fmt.Println(playerTwoName)
	//
	//
	GameBoard := o.InitBoard("brown")
	GameBoard.PrintCurrentBoard()
	//
	for winCondition == false {
		// on top of currMove, do currPlayer
		currMove := o.Move{}
		fmt.Printf("\n%v's Move!",GameBoard.Turn)
		if GameBoard.Turn == "brown" {
			currMove = Player1.GetMoveFromHuman()
		} else {
			currMove = Player2.GetMoveFromHuman()
		}
		//
		if o.IsLegalMove(currMove) {
			// should add method that modifies Player's team
			GameBoard = GameBoard.UpdateBoard(currMove)
			GameBoard.PrintCurrentBoard()
		}
	}
}
