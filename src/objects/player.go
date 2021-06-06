package objects

//import "fmt"

//
// 0,1,2
var PIECE_SIZE_LIMIT int = 2
//
// constructor for Player
type Player struct {
	name string
	color bool
	team []Piece
}
//
// 
func (p Player) generateTeam() []Piece {
	var result []Piece
	for i:= 0; i <= PIECE_SIZE_LIMIT; i++ {
		tempPiece := Piece{color: p.color, size: i}
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
	p1 := Player{name: name1, color: true}
	p2 := Player{name: name2, color: false}
	p1.team = p1.generateTeam()
	p2.team = p2.generateTeam()
	return p1,p2
}