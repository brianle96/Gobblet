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
	GameBoard := o.InitBoard()
	GameBoard.PrintCurrentBoard()
}