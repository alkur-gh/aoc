package main

import (
    "fmt"
    "math"
    "day15/matrix"
)

func Min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

type Point struct {
    i, j int
}

func GetValidNeighbors(visit matrix.Matrix, i, j int) []Point {
    points := []Point{}
    for _, d := range []Point{{-1, 0}, {0, -1}, {1, 0}, {0, 1}} {
        if visit.IsValidIndex(i + d.i, j + d.j) {
            points = append(points, Point{i + d.i, j + d.j})
        }
    }
    return points
}

func FindMinUnvisited(visit, dist matrix.Matrix) Point {
    var p Point
    found := false
    rows, cols := visit.Shape()
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if visit.Get(i, j) != 0 {
                continue
            }
            if !found || dist.Get(p.i, p.j) > dist.Get(i, j) {
                p = Point{i, j}
                found = true
            }
        }
    }
    return p
}

func ExpandMatrix(m matrix.Matrix) matrix.Matrix {
    mr, mc := m.Shape()
    res := matrix.NewMatrix(5 * mr, 5 * mc)
    rows, cols := res.Shape()

    for i := 0; i < mr; i++ {
        for j := 0; j < mc; j++ {
            res.Set(i, j, m.Get(i, j))
        }
    }

    for i := 0; i < mr; i++ {
        for j := mc; j < cols; j++ {
            v := res.Get(i, j - mc) % 9 + 1
            res.Set(i, j, v)
        }
    }

    for i := mr; i < rows; i++ {
        for j := 0; j < cols; j++ {
            v := res.Get(i - mr, j) % 9 + 1
            res.Set(i, j, v)
        }
    }

    return res
}

func main() {
    path := "./files/handout.txt"
    risks := ExpandMatrix(matrix.ReadMatrix(path))
//    risks.Print()

    rows, cols := risks.Shape()

    visit := matrix.NewMatrix(rows, cols)
    dist := matrix.NewMatrix(rows, cols)
    dist.Fill(math.MaxInt64)

    i, j := 0, 0
    dist.Set(0, 0, 0)

    for visit.Get(rows - 1, cols - 1) == 0 {
        v := dist.Get(i, j)
        neighbors := GetValidNeighbors(visit, i, j)
        for _, n := range neighbors {
            pos := v + risks.Get(n.i, n.j)
            if pos < dist.Get(n.i, n.j) {
                dist.Set(n.i, n.j, pos)
            }
        }
        visit.Set(i, j, 1)

        m := FindMinUnvisited(visit, dist)
        i, j = m.i, m.j
    }

    fmt.Println(dist.Get(rows - 1, cols - 1))
}
