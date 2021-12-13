package utils

func Check(e error) {
    if e != nil {
        panic(e)
    }
}

func MakeMatrix(rows, cols int) [][]int {
    matrix := make([][]int, rows)
    for i := 0; i < rows; i++ {
        matrix[i] = make([]int, cols)
    }
    return matrix
}
