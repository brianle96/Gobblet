package main

import (
	"fmt"
	o "github.com/objects"
)

func main() {
	Player1, Player2 := o.GeneratePair("Brian","Erin")
	//
	playerOneName := Player1.GetPlayerName()
	playerTwoName := Player2.GetPlayerName()
	fmt.Println(playerOneName)
	fmt.Println(playerTwoName)
	//
	//
	GameBoard := o.InitBoard("brown")
	GameBoard.PrintCurrentBoard()
	//
	currMove := Player1.GetMoveFromHuman()
	//
	if o.IsLegalMove(currMove) {
		GameBoard.Grid = GameBoard.UpdateBoard(currMove)
		GameBoard.PrintCurrentBoard()
	}
}