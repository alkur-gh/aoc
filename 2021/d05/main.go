package main

import (
    "fmt"
    "day5/vents"
)

func main() {
    path := "./tests/files/input.txt"
    rows, cols, segments := vents.ReadSegments(path)
    ventsMap := vents.NewVentsMap(rows, cols)
    for _, segment := range segments {
        ventsMap.AddSegment(segment)
    }
//    ventsMap.Print()
    fmt.Println(ventsMap.CountIntersections())
}
