package main

import (
    "fmt"
    "os"
    "bufio"
    "math"
)

const (
    EMPTY rune = '.'
)

var TARGET_POSITIONS = map[FigureType][]Position{
    //'A': {Position{2, 3}, Position{3, 3}},
    //'B': {Position{2, 5}, Position{3, 5}},
    //'C': {Position{2, 7}, Position{3, 7}},
    //'D': {Position{2, 9}, Position{3, 9}},
    'A': {Position{2, 3}, Position{3, 3}, Position{4, 3}, Position{5, 3}},
    'B': {Position{2, 5}, Position{3, 5}, Position{4, 5}, Position{5, 5}},
    'C': {Position{2, 7}, Position{3, 7}, Position{4, 7}, Position{5, 7}},
    'D': {Position{2, 9}, Position{3, 9}, Position{4, 9}, Position{5, 9}},
}

var MOVE_COST = map[FigureType]int {
    'A': 1, 'B': 10, 'C': 100, 'D': 1000,
}

type Field struct {
    v [][]rune
}

type Position struct {
    i, j int
}

type FigureType rune

type Figure struct {
    figureType FigureType
    position Position
}

type Move struct {
    figure Figure
    dest Position
}

func abs(a int) int {
    if a > 0 {
        return a
    } else {
        return -a
    }
}

func (m *Move) Cost() int {
    c := abs(m.figure.position.j - m.dest.j)
    c += abs(m.dest.i - 1)
    if m.figure.position.i > 1 {
        c += abs(m.figure.position.i - 1)
    }
    return c * MOVE_COST[m.figure.figureType]
}

func (f *Field) Figures() map[FigureType][]Figure {
    fs := make(map[FigureType][]Figure)
    for i := 0; i < len(f.v); i++ {
        for j := 0; j < len(f.v[i]); j++ {
            r := FigureType(f.v[i][j])
            if (r == 'A') || (r == 'B') || (r == 'C') || (r == 'D') {
                fs[r] = append(fs[r], Figure{r, Position{i, j}})
            }
        }
    }
    return fs
}

func (f *Field) Print() {
    for _, rs := range f.v {
        for _, r := range rs {
            fmt.Printf("%c", r)
        }
        fmt.Println()
    }
    fmt.Println()
}

func ReadField(path string) *Field {
    f, err := os.Open(path)
    if err != nil {
        panic(err)
    }
    defer f.Close()
    scanner := bufio.NewScanner(f)
    v := [][]rune{}
    for scanner.Scan() {
        v = append(v, []rune(scanner.Text()))
    }
    return &Field{v}
}

func (f *Figure) PossibleMoves(field *Field) []Move {
    ms := []Move{}
    tps := TARGET_POSITIONS[f.figureType]
    v := field.v

    if f.position.j == tps[0].j {
        flag := true
        for i := tps[len(tps) - 1].i; i > f.position.i; i-- {
            if v[i][f.position.j] != rune(f.figureType) {
                flag = false
            }
        }
        if flag {
            return ms
        }
    }

    for i := 1; i < f.position.i; i++ {
        if v[i][f.position.j] != EMPTY {
            return ms
        }
    }

    left, right := f.position.j - 1, f.position.j + 1
    for v[1][left] == EMPTY {
        left--
    }
    for v[1][right] == EMPTY {
        right++
    }

    if (left < tps[0].j) && (tps[0].j < right) {
        for k := len(tps) - 1; k >= 0; k-- {
            tp := tps[k]
            if v[tp.i][tp.j] == EMPTY {
                return append(ms, Move{*f, tp})
            }
            if v[tp.i][tp.j] != rune(f.figureType) {
                break
            }
        }
    }

    if f.position.i == 1 {
        return ms
    }

    for j := left + 1; j < right; j++ {
        if (j == 3) || (j == 5) || (j == 7) || (j == 9) {
            continue
        }
        ms = append(ms, Move{*f, Position{1, j}})
    }

    return ms
}

func (f *Field) PossibleMoves() []Move {
    ms := []Move{}
    for _, fs := range f.Figures() {
        for _, figure := range fs {
            ms = append(ms, figure.PossibleMoves(f)...)
        }
    }
    return ms
}

func (f *Field) Copy() *Field {
    v := [][]rune{}
    for _, rs := range f.v {
        xs := []rune{}
        xs = append(xs, rs...)
        v = append(v, xs)
    }
    return &Field{v}
}

func (f *Field) MakeMove(move Move) *Field {
    res := f.Copy()
    res.v[move.figure.position.i][move.figure.position.j] = EMPTY
    res.v[move.dest.i][move.dest.j] = rune(move.figure.figureType)
    return res
}

func (f *Field) IsComplete() bool {
    for t, tps := range TARGET_POSITIONS {
        for _, tp := range tps {
            if f.v[tp.i][tp.j] != rune(t) {
                return false
            }
        }
    }
    return true
}

func rec(f *Field, score int, trace []Move, thresh int) (int, bool, []Move) {
    moves := f.PossibleMoves()
    min := math.MaxInt64
    //f.Print()
    if score > thresh {
        return score, false, trace
    }

    if len(moves) == 0 {
        //fmt.Println(score)
        //fmt.Println("TERM")
        return score, f.IsComplete(), trace
    }
    min_trace := trace
    for _, move := range moves {
        score, complete, cand := rec(f.MakeMove(move), score + move.Cost(),
                                     append(trace, move), min)
        if complete && (score < min) {
            min = score
            min_trace = cand
        }
    }
    return min, min != math.MaxInt64, min_trace
}

func main() {
    path := "./files/input.txt"
    field := ReadField(path)
    score, _, trace := rec(field, 0, []Move{}, math.MaxInt64)
    fmt.Println(score)
    field.Print()
    acc := 0
    for _, move := range trace {
        acc += move.Cost()
        fmt.Println(acc)
        field = field.MakeMove(move)
        field.Print()
    }
    //fmt.Println(field.IsComplete())
    //field.Print()
    //for _, move := range field.PossibleMoves() {
    //    fmt.Println(move)
    //    fmt.Println(move.Cost())
    //    field.MakeMove(move).Print()
    //}
}
