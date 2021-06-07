package objects

type Move struct {
	Size int
	Row int
	Col int
	Color string
}
//
// should i allow a player to overlap their own piece
func IsLegalMove(move Move, gameBoard Board, currPlayer Player) bool {
	if !isInBounds(move) {
		return false
	} else if !isRightSize(move, gameBoard) {
		return false
	} else if !isInTeam(move, currPlayer) {
		return false
	}
	return true
}
//
//
func isInBounds(move Move) bool {
	result := true
	if move.Row < 0 || move.Row >= 3 {
		result = false
	}
	if move.Col < 0 || move.Col >= 3 {
		result = false
	}
	return result
}
//
//
func isRightSize(move Move, gameBoard Board) bool {
	result := true
	intendedPosition := gameBoard.Grid[move.Row][move.Col].Size
	if move.Size <= intendedPosition {
		result = false
	}
	return result
}
//
//
func isInTeam(move Move,currPlayer Player) bool {
	result := false
	tempPiece := Piece{Color: move.Color, Size: move.Size}
	for _, copy := range currPlayer.Team {
		if copy == tempPiece {
			result = true
		}
	}
	return result
}