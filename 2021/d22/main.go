package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
    "regexp"
    g "day22/graphics"
)

const (
    ON string = "on"
    OFF string = "off"
)

type RebootStep struct {
    action string
    x0, x1, y0, y1, z0, z1 int
}

func ReadRebootSteps(path string) []*RebootStep {
    f, err := os.Open(path)
    if err != nil {
        panic(err)
    }
    defer f.Close()
    scanner := bufio.NewScanner(f)
    steps := []*RebootStep{}
    r := regexp.MustCompile(
        `(?P<action>on|off) ` +
        `x=(?P<x0>-?\d+)\.\.(?P<x1>-?\d+),` +
        `y=(?P<y0>-?\d+)\.\.(?P<y1>-?\d+),` +
        `z=(?P<z0>-?\d+)\.\.(?P<z1>-?\d+)`)

    for scanner.Scan() {
        ms := r.FindStringSubmatch(scanner.Text())
        action := ms[1]
        x0, _ := strconv.Atoi(ms[2])
        x1, _ := strconv.Atoi(ms[3])
        y0, _ := strconv.Atoi(ms[4])
        y1, _ := strconv.Atoi(ms[5])
        z0, _ := strconv.Atoi(ms[6])
        z1, _ := strconv.Atoi(ms[7])
        step := &RebootStep{action, x0, x1, y0, y1, z0, z1}
        steps = append(steps, step)
    }

    return steps
}

func main() {
    //g.Demo()
    //fmt.Println("HELL")
    path := "./files/simple.txt"
    steps := ReadRebootSteps(path)
    rects := []*g.Rectangle{}
    offs := []*g.Rectangle{}
    runes := "xo-+*qwertyuip"

    for i, s := range steps {
        nr := g.MakeRectangle(
            s.x0, s.x1, s.y0, s.y1,
            rune(runes[i % len(runes)]),
        )
        next := []*g.Rectangle{}
        if s.action == ON {
            next = append(rects, nr)
        } else {
            //next = rects
            //offs = append(offs, nr)
            for _, r := range rects {
                next = append(next, r.Difference(nr)...)
            }
        }
        rects = next
        fmt.Println(s)
    }

    //union := []*g.Rectangle{rects[0]}
    //for _, r := range rects[1:] {
    //    next := []*g.Rectangle{}
    //    for _, t := range union {
    //        next = append(next, t.Union(r)...)
    //    }
    //    union = next
    //}

    fmt.Println(offs)

    cvs := g.MakeCanvas(g.MakeBorders(-50, 50, -50, 50, -50, 50))

    area := 0
    for _, r := range rects {
        area += r.Area()
        cvs.DrawRectangle(r)
    }

    for _, r := range offs {
        cvs.ClearRectangle(r)
    }
    cvs.Plot()
    fmt.Println(cvs.CountNonEmpty())
    //fmt.Println(area)
}
