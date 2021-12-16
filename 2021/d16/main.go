package main

import (
    "fmt"
    "day16/packet"
)

func main() {
    p := packet.ExpandHexFromFile("./tests/files/input.txt")
    fmt.Println(p.LiteralValue())
}
