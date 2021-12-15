package matrix

import (
    "fmt"
    "os"
    "bufio"
)

type Matrix [][]int

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func NewMatrix(rows, cols int) Matrix {
    matrix := make([][]int, rows)
    for i := 0; i < rows; i++ {
        matrix[i] = make([]int, cols)
    }
    return matrix
}

func ReadMatrix(path string) Matrix {
    f, err := os.Open(path)
    check(err)
    defer f.Close()
    scanner := bufio.NewScanner(f)
    var matrix [][]int
    for scanner.Scan() {
        var row []int
        for _, r := range scanner.Text() {
            row = append(row, int(r - '0'))
        }
        matrix = append(matrix, row)
    }
    return matrix
}

func (m Matrix) Print() {
    for _, row := range m {
        for _, v := range row {
            fmt.Printf("%1d", v)
        }
        fmt.Println()
    }
    fmt.Println()
}

func (m Matrix) Fill(filler int) {
    rows, cols := m.Shape()
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            m[i][j] = filler
        }
    }
}

func (m Matrix) Set(i, j int, value int) {
    m[i][j] = value
}

func (m Matrix) Get(i, j int) int {
    return m[i][j]
}

func (m Matrix) Shape() (int, int) {
    return len(m), len(m[0])
}

func (m Matrix) IsValidIndex(i, j int) bool {
    rows, cols := m.Shape()
    return i >= 0 && i < rows && j >= 0 && j < cols
}
