package main

import (
    "fmt"
    "day7/crabs"
)

func main() {
    path := "./tests/files/input.txt"
    positions := crabs.ReadPositions(path)
    pos, fuel := crabs.FindPositionMinimizingFuel(positions)
    fmt.Println(pos, fuel)
}
