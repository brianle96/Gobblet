package modules
//
// 1,2,3
var PIECE_SIZE_LIMIT int = 3
var NUM_PIECES int = 2
//
//
type Piece struct {
	Player string
	Size int
}
//
// 
func GenerateTeam(team string) []Piece {
	var result []Piece
	for i:= 1; i <= PIECE_SIZE_LIMIT; i++ {
		tempPiece := Piece{Player:team, Size: i}
		for j := 1; j <= NUM_PIECES; j++ {
			result = append(result, tempPiece)
		}
	}
	//fmt.Println(result)
	return result
}
//
//
func RemovePiece(gameBoard Board, currMove Move) Board{
	//
	if gameBoard.Turn {
		gameBoard.Pieces1 = RemovePieceFromTeam(gameBoard.Pieces1, currMove)
	} else {
		gameBoard.Pieces2 = RemovePieceFromTeam(gameBoard.Pieces2, currMove)
	}
	return gameBoard
}
func RemovePieceFromTeam (team []Piece, currMove Move) []Piece {
	if len(team) == 1 {
		return nil
	}
	tempPiece := Piece{Player: currMove.Player, Size: currMove.Size}
	for index, copy := range team {
		if copy == tempPiece {
			team = removeElementFromSlice(team, index)
			break
		}
	}
	return team
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