package main

import (
    "fmt"
    "os"
    "bufio"
    "sort"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

type Point struct {
    i, j int
}

func ReadHeightmap(path string) []string {
    f, err := os.Open(path)
    check(err)
    defer f.Close()
    scanner := bufio.NewScanner(f)
    var lines []string
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines
}

func getVisitablePoints(heightmap []string, i, j int) []Point {
    rows, cols := len(heightmap), len(heightmap[0])
    var points []Point
    if i > 0 {
        points = append(points, Point{i - 1, j})
    }
    if i < rows - 1 {
        points = append(points, Point{i + 1, j})
    }
    if j > 0 {
        points = append(points, Point{i, j - 1})
    }
    if j < cols - 1 {
        points = append(points, Point{i, j + 1})
    }
    return points
}

func isLowPoint(heightmap []string, i, j int) bool {
    h := heightmap[i][j]
    for _, p := range getVisitablePoints(heightmap, i, j) {
        other_height := heightmap[p.i][p.j]
        if h >= other_height {
            return false
        }
    }
    return true
}

func findLowPoints(heightmap []string) []Point {
    rows, cols := len(heightmap), len(heightmap[0])
    var points []Point
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if isLowPoint(heightmap, i, j) {
                points = append(points, Point{i, j})
            }
        }
    }
    return points
}

func visitBasin(heightmap []string, i, j int, visited [][]bool) []Point {
    h := heightmap[i][j]
    basin := []Point{Point{i, j}}
    visited[i][j] = true
    for _, p := range getVisitablePoints(heightmap, i, j) {
        other_height := heightmap[p.i][p.j]
        if !visited[p.i][p.j] && other_height != '9' && h < other_height {
            basin = append(basin, visitBasin(heightmap, p.i, p.j, visited)...)
        }
    }
    return basin
}

func prod(nums []int) int {
    res := 1
    for _, num := range nums {
        res *= num
    }
    return res
}

func makeVisitedMatrix(rows, cols int) [][]bool {
    var visited [][]bool
    for i := 0; i < rows; i++ {
        visited = append(visited, make([]bool, cols))
    }
    return visited
}

func GetBasinScore(heightmap []string) int {
    points := findLowPoints(heightmap)
    rows, cols := len(heightmap), len(heightmap[0])

    basinSizes := make([]int, 0)
    for _, p := range points {
        visited := makeVisitedMatrix(rows, cols)
        basin := visitBasin(heightmap, p.i, p.j, visited)
        basinSizes = append(basinSizes, len(basin))
    }

    if len(basinSizes) > 3 {
        sort.Ints(basinSizes)
        basinSizes = basinSizes[len(basinSizes) - 3:]
    }

    return prod(basinSizes)
}

func main() {
    path := "./files/handout.txt"
    heightmap := ReadHeightmap(path)
    score := GetBasinScore(heightmap)
    fmt.Println(score)
}
