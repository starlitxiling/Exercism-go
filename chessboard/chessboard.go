package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	// panic("Please implement CountInFile()")
	count := 0
	for k, v := range cb {
		if k == file {
			for _, occupied := range v {
				if occupied {
					count += 1
				}
			}
		}
	}
	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	// panic("Please implement CountInRank()")
	count := 0
	for _, v := range cb {
		for i, occupied := range v {
			if i+1 == rank {
				if occupied {
					count += 1
				}
			}
		}
	}
	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	// panic("Please implement CountAll()")
	for _, v := range cb {
		return len(v) * len(cb)
	}
	return 0
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	// panic("Please implement CountOccupied()")
	count := 0
	for _, v := range cb {
		for _, occupied := range v {
			if occupied {
				count += 1
			}
		}
	}
	return count
}
