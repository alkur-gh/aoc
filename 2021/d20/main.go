package main

import (
    "fmt"
    "os"
    "bufio"
)

type Grid [][]byte

func (g Grid) Shape() (int, int) {
    return len(g), len(g[0])
}

func MakeGrid(rows, cols int) Grid {
    g := [][]byte{}
    for i := 0; i < rows; i++ {
        g = append(g, make([]byte, cols))
    }
    return g
}

func (g Grid) Print() {
    fmt.Printf("[%d, %d]:\n", len(g), len(g[0]))
    for _, row := range g {
        for _, elem := range row {
            if elem == 1 {
                fmt.Print("#")
            } else {
                fmt.Print(".")
            }
        }
        fmt.Println()
    }
    fmt.Println()
}

func ReadInput(path string) ([]byte, Grid) {
    f, err := os.Open(path)
    if err != nil {
        panic(err)
    }
    defer f.Close()
    scanner := bufio.NewScanner(f)
    scanner.Scan()

    algo := []byte{}
    for _, r := range scanner.Text() {
        if r == '#' {
            algo = append(algo, 1)
        } else {
            algo = append(algo, 0)
        }
    }

    scanner.Scan()

    grid := [][]byte{}
    for scanner.Scan() {
        row := []byte{}
        for _, r := range scanner.Text() {
            if r == '#' {
                row = append(row, 1)
            } else {
                row = append(row, 0)
            }
        }
        grid = append(grid, row)
    }

    return algo, grid
}

func (g Grid) IsValidIndex(i, j int) bool {
    return (i >= 0) && (i < len(g)) && (j >= 0) && (j < len(g[0]))
}

func (g Grid) Enhance(algo []byte, border int) Grid {
    neighbors := []struct {
        di, dj int
    }{
        {1, 1}, {1, 0}, {1, -1},
        {0, 1}, {0, 0}, {0, -1},
        {-1, 1}, {-1, 0}, {-1, -1},
    }
    grows, gcols := g.Shape()
    rows, cols := grows + 2, gcols + 2
    res := MakeGrid(rows, cols)
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            v := 0
            for k, nb := range neighbors {
                gi, gj := i + nb.di - 1, j + nb.dj - 1
                if g.IsValidIndex(gi, gj) {
                    v |= (int(g[gi][gj]) << k)
                } else {
                    v |= (border << k)
                }
            }
            res[i][j] = algo[v]
        }
    }
    return res
}

func (g Grid) CountLit() int {
    count := 0
    for _, row := range g {
        for _, elem := range row {
            count += int(elem)
        }
    }
    return count
}

func (g Grid) EnhanceMultipleTimes(algo []byte, times int) Grid {
    enh := g
    border := 0
    for i := 0; i < times; i++ {
        enh = enh.Enhance(algo, border)
        if border == 0 {
            border = int(algo[0])
        } else {
            border = int(algo[len(algo) - 1])
        }
    }
    return enh
}

func main() {
    path := "./files/input.txt"
    times := 50
    algo, grid := ReadInput(path)
    enh := grid.EnhanceMultipleTimes(algo, times)
    enh.Print()
    fmt.Println(enh.CountLit())
}
