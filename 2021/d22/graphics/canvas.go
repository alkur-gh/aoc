package graphics

import "fmt"

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

func (c *Canvas) DrawRectangle(rs ...*Rectangle) {
    xo, yo := c.length / 2, c.height / 2
    for _, r := range rs {
        if r == nil {
            continue
        }
        for x := r.x0; x <= r.x1; x++ {
            for y := r.y0; y <= r.y1; y++ {
                c.v[y + yo][x + xo] = r.r
            }
        }
    }
}

func (c *Canvas) ClearRectangle(rs ...*Rectangle) {
    xo, yo := c.length / 2, c.height / 2
    for _, r := range rs {
        if r == nil {
            continue
        }
        for x := r.x0; x <= r.x1; x++ {
            for y := r.y0; y <= r.y1; y++ {
                c.v[y + yo][x + xo] = 0
            }
        }
    }
}

func (c *Canvas) DrawLine(ls ...*Line) {
    xo, yo := c.length / 2, c.height / 2
    for _, l := range ls {
        if l == nil {
            continue
        }
        for x := l.x0; x <= l.x1; x++ {
            c.v[yo][x + xo] = l.r
        }
    }
}

func (c *Canvas) CountNonEmpty() int {
    count := 0
    for _, rs := range c.v {
        for _, r := range rs {
            if r != 0 {
                count++
            }
        }
    }
    return count
}
