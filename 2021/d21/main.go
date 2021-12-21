package main

import (
    "fmt"
)

const THRESHOLD = 21
var ROLLS = map[int]int {3: 1, 4: 3, 5: 6, 6: 7, 7: 6, 8: 3, 9: 1}

type Player struct {
    position, score int
}

func (p Player) Move(n int) Player {
    v := (p.position + n - 1) % 10 + 1
    return Player{v, p.score + v}
}

func rec(p1, p2 Player, first bool) (int, int) {
    if p1.score >= THRESHOLD {
        return 1, 0
    } else if p2.score >= THRESHOLD {
        return 0, 1
    } else if first {
        ls, rs := 0, 0
        for k, v := range ROLLS {
            l, r := rec(p1.Move(k), p2, false)
            ls += l * v
            rs += r * v
        }
        return ls, rs
    } else {
        ls, rs := 0, 0
        for k, v := range ROLLS {
            l, r := rec(p1, p2.Move(k), true)
            ls += l * v
            rs += r * v
        }
        return ls, rs
    }
}

func main() {
    p1 := Player{3, 0}
    p2 := Player{7, 0}
    fmt.Println(rec(p1, p2, true))
}
