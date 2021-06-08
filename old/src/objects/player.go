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
var NUM_PIECES int = 2
//
// constructor for Player
type Player struct {
	Color string
	Team []Piece
	move Move
}
//
// 
func (p Player) generateTeam() []Piece {
	var result []Piece
	for i:= 1; i <= PIECE_SIZE_LIMIT; i++ {
		tempPiece := Piece{Color: p.Color, Size: i}
		for j := 1; j <= NUM_PIECES; j++ {
			result = append(result, tempPiece)
		}
	}
	//fmt.Println(result)
	return result
}
//
//
func (p Player) GetMoveFromHuman() Move {
	fmt.Println("\n\nEnter your move (row,col,size):")

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
	tempMove := Move{Size: currSize, Row: rowMove, Col: colMove, Color: p.Color}
	return tempMove
}
//
//
func removeElementFromSlice(team []Piece, index int) []Piece {
	// Remove the element at index i from a.
	team[index] = team[len(team)-1] // Copy last element to index i.
	team[len(team)-1] = Piece{}  // Erase last element (write zero value).
	team = team[:len(team)-1]   // Truncate slice.
	return team
}
//
//
func (p Player) RemovePiece(move Move) []Piece {
	if len(p.Team) == 1 {
		return nil
	}
	tempPiece := Piece{Color: move.Color, Size: move.Size}
	for index, copy := range p.Team {
		if copy == tempPiece {
			p.Team = removeElementFromSlice(p.Team, index)
			break
		}
	}
	return p.Team
}
//
// return a pair of players
func GeneratePair(name1 string, name2 string) (Player, Player) {
	p1 := Player{Color: "brown"}
	p2 := Player{Color: "green"}
	p1.Team = p1.generateTeam()
	p2.Team = p2.generateTeam()
	return p1,p2
}