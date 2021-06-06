package main

import (
	"fmt"
	o "github.com/objects"
)

func main() {
	player1, player2 := o.GeneratePair("Brian","Erin")
	//
	playerOneName := player1.GetPlayerName()
	playerTwoName := player2.GetPlayerName()
	fmt.Println(playerOneName)
	fmt.Println(playerTwoName)
	//fmt.Println(player1.GetPlayerTeam())
	//fmt.Println(player2.GetPlayerTeam())
}