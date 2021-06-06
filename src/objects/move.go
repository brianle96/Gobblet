package objects

type Move struct {
	Size int
	Row int
	Col int
	Color string
}
//
//
func IsLegalMove(move Move) bool {
	if !isInBounds(move) {
		return false
	}
	return true
}
//
//
func isInBounds(move Move) bool {
	result := true
	if move.Row < 1 || move.Row > 3 {
		result = false
	}
	if move.Col < 1 || move.Col > 3 {
		result = false
	}
	return result
}