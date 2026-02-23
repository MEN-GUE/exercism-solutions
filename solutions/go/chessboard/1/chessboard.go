package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File


// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
    used := 0
    // Validate file input
    _, ok := cb[file]
    if ok == false {
        return 0
    }

    // Look for occupied spaces
	for _, v := range cb[file] {
        if v == true {
            used++
        }
    }
    return used
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	used := 0
	// Validate rank input
    if rank <= 0 || rank > 8 {
        return 0
    }

    // Look for occupied spaces
    for _, v := range cb {
        if v[rank -1] == true {
            used ++
        }
    }
    return used
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
    files := 0
    squaresPerFile := 0
    //Count files
	for range cb {
        files++
    }

    // Count files squares bia the first one
    for range cb["A"] {
        squaresPerFile++
    }

    return files * squaresPerFile
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
    squares := 0
    for v,_ := range cb {
        squares += CountInFile(cb, v)
    }
    return squares
}
