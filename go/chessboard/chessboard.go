package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	var occupied int
	if val, ok := cb[file]; ok {
		for _, i := range val {
			if i {
				occupied++
			}
		}
		return occupied
	}
	return 0
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	var occupied int
	if rank > 0 && rank < 9 {
		for _, val := range cb {
			if val[rank-1] {
				occupied++
			}
		}
		return occupied
	}
	return 0
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	var total int
	for _, val := range cb {
		total += len(val)
	}
	return total
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var occupied int
	for _, i := range cb {
		for _, j := range i {
			if j {
				occupied++
			}
		}
	}
	return occupied
}
