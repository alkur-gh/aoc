package main

import (
    "fmt"
    "os"
    "bufio"
)

const (
    EMPTY rune = '.'
)

var TARGET_POSITIONS = map[FigureType][]Position{
    'A': {Position{2, 3}, Position{3, 3}},
    'B': {Position{2, 5}, Position{3, 5}},
    'C': {Position{2, 7}, Position{3, 7}},
    'D': {Position{2, 9}, Position{3, 9}},
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

    if ((f.position == tps[0]) &&
        (v[tps[1].i][tps[1].j] == rune(f.figureType))) ||
       (f.position == tps[1]) {
        return ms
    }

    if (f.position.i == 3) &&
       (v[f.position.i - 1][f.position.j] != EMPTY) {
        return ms 
    }


    left, right := f.position.j - 1, f.position.j + 1
    for v[1][left] == EMPTY {
        left--
    }
    for v[1][right] == EMPTY {
        right++
    }

    //if f.position.i == 1 {
    //    high, low := tps[0], tps[1]
    //    if (v[low.i][low.j] == rune(f.figureType)) &&
    //       (v[high.i][high.j] == EMPTY) {
    //        ms = append(ms, Move{*f, high})
    //    }
    //    if (v[low.i][low.j] == EMPTY) &&
    //       (v[high.i][high.j] == EMPTY) {
    //        ms = append(ms, Move{*f, low})
    //    }
    //    return ms
    //}

    high, low := tps[0], tps[1]
    if (left < low.j) && (low.j < right) {
        if v[low.i][low.j] == EMPTY {
            return append(ms, Move{*f, low})
        }
        if (v[low.i][low.j] == rune(f.figureType)) && (v[high.i][high.j] == EMPTY) {
            return append(ms, Move{*f, high})
        }
    }

    for j := left + 1; j < right; j++ {
        if f.position.j == j {
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
        xs:= []rune{}
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

func rec(f *Field) {
    f.Print()
    moves := f.PossibleMoves()
    if len(moves) == 0 {
        fmt.Println("TERM")
        return
    }
    for _, move := range moves {
        rec(f.MakeMove(move))
    }
}

func main() {
    path := "./files/test.txt"
    field := ReadField(path)
    rec(field)
//    field.Print()
//    for _, move := range field.PossibleMoves() {
//        fmt.Println(move)
//        field.MakeMove(move).Print()
//    }
}
