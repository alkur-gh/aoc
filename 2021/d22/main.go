package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
    "regexp"
)

const (
    ON string = "on"
    OFF string = "off"
    XMIN int = -50
    XMAX int = 50
    YMIN int = -50
    YMAX int = 50
    ZMIN int = -50
    ZMAX int = 50
)

type Borders struct {
    xmin, xmax, ymin, ymax, zmin, zmax int
}

type Grid [][][]bool

func MakeGrid(m, n, t int) Grid {
    g3 := [][][]bool{}
    for i := 0; i < m; i++ {
        g2 := [][]bool{}
        for j := 0; j < n; j++ {
            g2 = append(g2, make([]bool, t))
        }
        g3 = append(g3, g2)
    }
    return g3
}

func (g3 Grid) CountTurnedOn() int {
    count := 0
    for _, g2 := range g3 {
        for _, g1 := range g2 {
            for _, elem := range g1 {
                if elem {
                    count++
                }
            }
        }
    }
    return count
}

type RebootStep struct {
    mode string
    xl, xr, yl, yr, zl, zr int
}

func ReadInput(path string) ([]RebootStep, Borders) {
    f, err := os.Open(path)
    if err != nil {
        panic(err)
    }
    defer f.Close()
    scanner := bufio.NewScanner(f)
    steps := []RebootStep{}
    r := regexp.MustCompile(
        `(?P<mode>on|off) ` +
        `x=(?P<xl>-?\d+)\.\.(?P<xr>-?\d+),` +
        `y=(?P<yl>-?\d+)\.\.(?P<yr>-?\d+),` +
        `z=(?P<zl>-?\d+)\.\.(?P<zr>-?\d+)`)

    xmin, xmax, ymin, ymax, zmin, zmax := 0, 0, 0, 0, 0, 0

    for scanner.Scan() {
        ms := r.FindStringSubmatch(scanner.Text())
        mode := ms[1]
        xl, _ := strconv.Atoi(ms[2])
        xr, _ := strconv.Atoi(ms[3])
        yl, _ := strconv.Atoi(ms[4])
        yr, _ := strconv.Atoi(ms[5])
        zl, _ := strconv.Atoi(ms[6])
        zr, _ := strconv.Atoi(ms[7])
        xmin = min(xmin, xl)
        xmax = max(xmax, xr)
        ymin = min(ymin, yl)
        ymax = max(ymax, yr)
        zmin = min(zmin, zl)
        zmax = max(zmax, zr)
        step := RebootStep{mode, xl, xr, yl, yr, zl, zr}
        steps = append(steps, step)
    }

    return steps, Borders{xmin, xmax, ymin, ymax, zmin, zmax}
}

func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

type Rectangle struct {
    x0, y0, x, y int
    r rune
}

type Canvas struct {
    length, height int
    v [][]rune
}

func MakeCanvas(b Borders) *Canvas {
    length := b.xmax - b.xmin + 1
    height := b.ymax - b.ymin + 1
    v := [][]rune{}
    for i := 0; i < height; i++ {
        v = append(v, make([]rune, height))
    }
    return &Canvas{length, height, v}
}

func (c *Canvas) Plot() {
    for _, rs := range c.v {
        for _, r := range rs {
            if r == 0 {
                fmt.Print(" ")
            } else {
                fmt.Printf("%c", r)
            }
        }
        fmt.Println()
    }
    fmt.Println()
}

func (c *Canvas) DrawRectangle(r *Rectangle) {
    xo, yo:= c.length / 2, c.height / 2
    for x := r.x0; x <= r.x; x++ {
        for y := r.y0; y <= r.y; y++ {
            c.v[y + yo][x + xo] = r.r
        }
    }
}

func (r1 *Rectangle) Intersection(r2 *Rectangle) *Rectangle {
    var left, right *Rectangle
    r := Rectangle{0, 0, 0, 0, '-'}
    if r1.x0 <= r2.x0 {
        left, right = r1, r2
    } else {
        left, right = r2, r1
    }
    if left.x < right.x0 {
        return nil
    }
    r.x0 = right.x0
    r.x = left.x
    return &r
}


func main() {
//    path := "./files/handout.txt"
//    steps, b := ReadInput(path)
    //b := Borders{XMIN, XMAX, YMIN, YMAX, ZMIN, ZMAX}
    b := Borders{-20, 20, -20, 20, -20, 20}
    r1 := &Rectangle{-10, -5, 10, 5, 'x'}
    r2 := &Rectangle{0, -4, 15, 0, 'o'}
    r := r1.Intersection(r2)
    cvs := MakeCanvas(b)
    cvs.DrawRectangle(r1)
    cvs.DrawRectangle(r2)
    cvs.DrawRectangle(r)
    cvs.Plot()
    //g := MakeGrid(b.xmax - b.xmin + 1, b.ymax - b.ymin + 1, b.zmax - b.zmin + 1)
    //fmt.Println(b)
    //for _, s := range steps {
    //    v := (s.mode == ON)
    //    for x := max(b.xmin, s.xl); x <= min(b.xmax, s.xr); x++ {
    //        for y := max(b.ymin, s.yl); y <= min(b.ymax, s.yr); y++ {
    //            for z := max(b.zmin, s.zl); z <= min(b.zmax, s.zr); z++ {
    //                g[x - b.xmin][y - b.ymin][z - b.zmin] = v
    //            }
    //        }
    //    }
    //}
    //fmt.Println(g.CountTurnedOn())
}
