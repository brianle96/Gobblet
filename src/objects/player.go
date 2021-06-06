package objects

//import "fmt"

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
	for i:= 0; i <= PIECE_SIZE_LIMIT; i++ {
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
// return a pair of players
func GeneratePair(name1 string, name2 string) (Player, Player) {
	p1 := Player{name: name1, color: "red"}
	p2 := Player{name: name2, color: "blue"}
	p1.team = p1.generateTeam()
	p2.team = p2.generateTeam()
	return p1,p2
}