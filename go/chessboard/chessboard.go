package chessboard

import "fmt"

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	counter := 0
	_, found := cb[file]
	if found {
		for i := 0; i < len(cb[file]); i++ {
			if cb[file][i] {
				counter++
			}
		}
	}
	return counter
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	counter := 0
	fmt.Println(rank)
	fmt.Println(cb)
	if rank >= 1 && rank <= 8 {
		for _, v := range cb {
			if v[rank-1] {
				counter++
			}
		}
	}
	return counter
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	counter := 0
	for range cb {
		counter++
	}
	return counter * counter
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	counter := 0
	for _, v := range cb {
		for i := 0; i < len(v); i++ {
			if v[i] {
				counter++
			}
		}
	}
	return counter
}
