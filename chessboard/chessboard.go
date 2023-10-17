package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	row := cb[file]
	sum := 0

	for _, x := range row {
		if x {
			sum += 1
		}
	}

	return sum
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank <= 0 || rank > 8 {
		return 0
	}

	sum := 0

	for _, x := range cb {
		if x[rank-1] {
			sum += 1
		}
	}

	return sum
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	sum := 0

	for _, x := range cb {
		sum += len(x)
	}

	return sum
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	sum := 0

	for k := range cb {
		sum += CountInFile(cb, k)
	}

	return sum
}
