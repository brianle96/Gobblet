package main

import (
	"fmt"
	"github.com/objects"
)

func main() {
	PlayerOne := objects.Player{Name: "Brian"}
	playerOneName := PlayerOne.GetPlayerName()
	fmt.Println(playerOneName)
	//
	PlayerTwo := objects.Player{Name: "Erin"}
	playerTwoName := PlayerTwo.GetPlayerName()
	fmt.Println(playerTwoName)
}