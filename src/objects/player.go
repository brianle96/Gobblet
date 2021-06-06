package objects

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"log"
)

//
// 1,2,3
var PIECE_SIZE_LIMIT int = 3
//
// constructor for Player
type Player struct {
	name string
	color string
	team []Piece
}
//
// 
func (p Player) generateTeam() []Piece {
	var result []Piece
	for i:= 1; i <= PIECE_SIZE_LIMIT; i++ {
		tempPiece := Piece{Color: p.color, Size: i}
		result = append(result, tempPiece)
		result = append(result, tempPiece)
	}
	//fmt.Println(result)
	return result
}
//
// return the player name
func (p Player) GetPlayerName() string {
	return p.name
}
//
// return the player team
func (p Player) GetPlayerTeam() []Piece {
	return p.team
}
//
//
func (p Player) GetMoveFromHuman() Move {
	fmt.Println("\nEnter your move (x,y,size):")

	var reader = bufio.NewReader(os.Stdin)
	message, _ := reader.ReadString('\n')
	
	re := regexp.MustCompile("[0-9]+")
	
	moveSlice := re.FindAllString(message, -1)

	rowMove, errRow := strconv.Atoi(moveSlice[0])
	// If an error was returned, print it to the console and
    // exit the program.
    if errRow != nil {
        log.Fatal(errRow)
    }
	colMove, errCol := strconv.Atoi(moveSlice[1])
	// If an error was returned, print it to the console and
    // exit the program.
    if errCol != nil {
        log.Fatal(errCol)
    }
	//
	currSize, errSize := strconv.Atoi(moveSlice[2])
	// If an error was returned, print it to the console and
    // exit the program.
    if errSize != nil {
        log.Fatal(errCol)
    }
	//
	tempMove := Move{Size: currSize, Row: rowMove, Col: colMove, Color: p.color}
	return tempMove
}
//
// return a pair of players
func GeneratePair(name1 string, name2 string) (Player, Player) {
	p1 := Player{name: name1, color: "brown"}
	p2 := Player{name: name2, color: "green"}
	p1.team = p1.generateTeam()
	p2.team = p2.generateTeam()
	return p1,p2
}