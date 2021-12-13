package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
    "day13/utils"
)

const X_AXIS = "x"
const Y_AXIS = "y"

type Paper [][]int

type Point struct {
    x, y int
}

type FoldCommand struct {
    axis string
    value int
}


func (paper Paper) shape() (int, int) {
    return len(paper), len(paper[0])
}

func (paper Paper) printit() {
    rows, cols := paper.shape()
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if paper[i][j] > 0 {
                fmt.Print("*")
            } else {
                fmt.Print(".")
            }
        }
        fmt.Println()
    }
}

func parsePoint(line string) Point {
    parts := strings.Split(line, ",")
    x, err := strconv.Atoi(parts[0])
    utils.Check(err)
    y, err := strconv.Atoi(parts[1])
    utils.Check(err)
    return Point{x, y}
}

func parseFoldCommand(line string) FoldCommand {
    parts := strings.Split(strings.Split(line, " ")[2], "=")
    axis := parts[0]
    value, err := strconv.Atoi(parts[1])
    utils.Check(err)
    return FoldCommand{axis, value}
}

func ReadPaperPointAndFoldPositions(path string) (Paper, []FoldCommand) {
    f, err := os.Open(path)
    utils.Check(err)
    defer f.Close()
    scanner := bufio.NewScanner(f)

    var points []Point
    maxX, maxY := 0, 0
    for scanner.Scan() {
        line := scanner.Text()
        if line == "" {
            break
        }
        point := parsePoint(line)
        points = append(points, point)
        if point.x >= maxX { maxX = point.x }
        if point.y >= maxY { maxY = point.y }
    }

    paper := utils.MakeMatrix(maxY + 1, maxX + 1)
    for _, p := range points {
        paper[p.y][p.x] = 1
    }

    var folds []FoldCommand
    for scanner.Scan() {
        folds = append(folds, parseFoldCommand(scanner.Text()))
    }

    return paper, folds
}

func (paper Paper) applyFoldCommand(fold FoldCommand) Paper {
    rows, cols := paper.shape()
    var folded Paper
    if fold.axis == X_AXIS {
        folded = utils.MakeMatrix(rows, fold.value)
        for i := 0; i < rows; i++ {
            for j := 0; j < fold.value; j++ {
                folded[i][j] = paper[i][j] + paper[i][cols - 1 - j]
            }
        }
    } else {
        folded = utils.MakeMatrix(fold.value, cols)
        for i := 0; i < fold.value; i++ {
            for j := 0; j < cols; j++ {
                folded[i][j] = paper[i][j] + paper[rows - 1 - i][j]
            }
        }
    }
    return folded
}

func (paper Paper) countVisible() int {
    count := 0
    for _, row := range paper {
        for _, elem := range row {
            if elem > 0 {
                count++
            }
        }
    }
    return count
}

func main() {
    path := "./files/input.txt"
    paper, folds := ReadPaperPointAndFoldPositions(path)
    for _, fold := range folds {
        paper = paper.applyFoldCommand(fold)
    }
    paper.printit()
}
