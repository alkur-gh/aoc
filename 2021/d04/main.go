package main

import (
    "fmt"
    "day4/bingo"
)

func main() {
    path := "./tests/files/input.txt"
    size := 5
    numbers, boards := bingo.ReadNumbersAndBoardsFromFile(path, size)
    minSteps, minStepsScore, maxSteps, maxStepsScore :=
        bingo.GetBestAndWorstBoards(numbers, boards)

    fmt.Printf("%d: %d\n", minSteps, minStepsScore)
    fmt.Printf("%d: %d\n", maxSteps, maxStepsScore)
}
