package main

import (
    "fmt"
)

type Point struct {
    x, y int
}

func (p Point) String() string {
    return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

type Rectangle struct {
    start, end Point
}

func (r Rectangle) String() string {
    return fmt.Sprintf("{%s,%s}", r.start, r.end)
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func abs(a int) int {
    if a < 0 {
        return -a
    } else {
        return a
    }
}

func (r Rectangle) Contains(p Point) bool {
    return (p.x >= r.start.x) && (p.x <= r.end.x) &&
           (p.y >= r.start.y) && (p.y <= r.end.y)
}

func (r Rectangle) CanBeHit(p Point, v Point) bool {
    for (p.x <= r.end.x) && ((v.y >= 0) || (p.y >= r.start.y)) {
        if r.Contains(p) {
            return true
        }
        p = Point{p.x + v.x, p.y + v.y}
        v = Point{max(v.x - 1, 0), v.y - 1}
    }
    return false
}

func (r Rectangle) CountHittingVelocities(p Point) int {
    count := 0
    ystart, yend := min(r.start.y, 0), max(abs(r.start.y), abs(r.end.y))
    for x := 0; x <= r.end.x; x++ {
        for y := ystart; y <= yend; y++ {
            if r.CanBeHit(p, Point{x, y}) {
                count++
            }
        }
    }
    return count
}

func main() {
    r := Rectangle{Point{20,-10}, Point{30,-5}}
    //r := Rectangle{Point{10, 10}, Point{20, 20}}
    fmt.Println(r.CountHittingVelocities(Point{0, 0}))
}
