package main

import (
    "fmt"
    lf "day6/lanternfish"
)

func main() {
    path := "./tests/files/input.txt"
    state := lf.ReadInitialState(path)
    days := 256
    for i := 0; i < days; i++ {
        state.Update()
    }
    fmt.Println(state.CountFish())
}
