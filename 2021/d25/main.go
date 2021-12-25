package main

import (
    "fmt"
    "os"
    "bufio"
)

const (
    EAST rune = '>'
    SOUTH rune = 'v'
    EMPTY rune = '.'
)

type State struct {
    m [][]rune
}

func (s *State) Shape() (int, int) {
    return len(s.m), len(s.m[0])
}

func (s *State) Print() {
    for _, row := range s.m {
        for _, r := range row {
            fmt.Printf("%c", r)
        }
        fmt.Println()
    }
    fmt.Println()
}

func (s *State) Next() (*State, bool) {
    rows, cols := s.Shape()
    m := make([][]rune, rows)
    for i := 0; i < rows; i++ {
        m[i] = make([]rune, cols)
        for j := 0; j < cols; j++ {
            m[i][j] = EMPTY
        }
    }
    moved := false
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if s.m[i][j] == EAST {
                nj := (j + 1) % cols
                if s.m[i][nj] == EMPTY {
                    moved = true
                    m[i][nj] = EAST
                } else {
                    m[i][j] = EAST
                }
            }
        }
    }
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if s.m[i][j] == SOUTH {
                ni := (i + 1) % rows
                if (m[ni][j] == EMPTY) && (s.m[ni][j] != SOUTH) {
                //if (s.m[ni][j] == EMPTY) || ((m[ni][j] == EAST) && (s.m[ni][j] != EAST)) {
                    moved = true
                    m[ni][j] = SOUTH
                } else {
                    m[i][j] = SOUTH
                }
            }
        }
    }
    return &State{m}, moved
}

func ReadState(path string) *State {
    f, err := os.Open(path)
    if err != nil { panic(err) }
    defer f.Close()
    scanner := bufio.NewScanner(f)
    m := [][]rune{}
    for scanner.Scan() {
        m = append(m, []rune(scanner.Text()))
    }
    return &State{m}
}

func main() {
    path := "./files/input.txt"
    state := ReadState(path)
    //state.Print()
    //s, flag := state.Next()
    //fmt.Println(flag)
    //s.Print()
    moved := true
    steps := 0
    for moved {
        steps++
        fmt.Println(steps)
        state, moved = state.Next()
    }
}
