package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	fileVal, fileExists := cb[file]
	if !fileExists {
		return 0
	}
	var total int
	for _, fileOccp := range fileVal {
		if fileOccp {
			total++
		}
	}
	return total
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
		return 0
	}
	var total int
	for _, rnk := range cb {
		if rnk[rank-1] {
			total++
		}
	}
	return total
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	var total int
	for _, file := range cb {
		for range file {
			total++
		}
	}
	return total
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var total int
	for _, file := range cb {
		for _, sq := range file {
			if sq {
				total++
			}
		}
	}
	return total
}
