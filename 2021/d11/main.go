package main

import (
    "fmt"
    "os"
    "bufio"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func makeMatrix(rows, cols int) [][]int {
    var matrix [][]int
    for i := 0; i < rows; i++ {
        matrix = append(matrix, make([]int, cols))
    }
    return matrix
}

func printMatrix(matrix [][]int) {
    for _, row := range matrix {
        for _, elem := range row {
            fmt.Print(elem, " ")
        }
        fmt.Println()
    }
}

func ReadEnergyLevels(path string) [][]int {
    f, err := os.Open(path)
    check(err)
    defer f.Close()
    scanner := bufio.NewScanner(f)
    var levels [][]int
    for scanner.Scan() {
        var nums []int
        for _, num := range scanner.Text() {
            nums = append(nums, int(num - '0'))
        }
        levels = append(levels, nums)
    }
    return levels
}

func MakeStep(levels [][]int) ([][]int, int) {
    rows, cols := len(levels), len(levels[0])
    flashed := makeMatrix(rows, cols)
    updated := makeMatrix(rows, cols)

    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            updated[i][j] = levels[i][j] + 1
        }
    }

    someFlashed := true
    for someFlashed {
        someFlashed = false
        for i := 0; i < rows; i++ {
            for j := 0; j < cols; j++ {
                if flashed[i][j] != 0 || updated[i][j] <= 9 {
                    continue
                }

                flashed[i][j] = 1
                updated[i][j] = 0
                someFlashed = true
                for di := -1; di < 2; di++ {
                    for dj := -1; dj < 2; dj++ {
                        p, t := i + di, j + dj
                        if p >= 0 && p < rows && t >= 0 && t < cols {
                            updated[p][t]++
                        }
                    }
                }
            }
        }
    }

    flashedCount := 0
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if flashed[i][j] != 0 {
                flashedCount++
                updated[i][j] = 0
            }
        }
    }

    return updated, flashedCount
}

func FindStepWhenAllFlashed(levels [][]int) int {
    size := len(levels) * len(levels[0])
    step := 0
    for {
        step++
        updated, flashed := MakeStep(levels)
        if flashed == size {
            break
        }
        levels = updated
    }
    return step
}

func main() {
    path := "./files/input.txt"
    levels := ReadEnergyLevels(path)
    fmt.Println(FindStepWhenAllFlashed(levels))
}
