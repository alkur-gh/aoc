package graphics

import "fmt"

func foo() {
    fmt.Println("FUCKING REQUIRED IMPORTS")
}

func (r *Rectangle) Points() (int, int, int, int) {
    return r.x0, r.x1, r.y0, r.y1
}

func (r *Rectangle) Area() int {
    return (r.x1 - r.x0 + 1) * (r.y1 - r.y0 + 1)
}

func (r *Rectangle) IsValid() bool {
    return (r.x0 <= r.x1) && (r.y0 <= r.y1)
}

func (r1 *Rectangle) Intersection(r2 *Rectangle) *Rectangle {
    var left, right *Rectangle
    r := Rectangle{0, 0, 0, 0, '-'}
    if r1.x0 <= r2.x0 {
        left, right = r1, r2
    } else {
        left, right = r2, r1
    }
    if left.x1 < right.x0 {
        return nil
    }
    r.x0 = right.x0
    r.x1 = min(left.x1, right.x1)

    if r1.y0 <= r2.y0 {
        left, right = r1, r2
    } else {
        left, right = r2, r1
    }
    if left.y1 < right.y0 {
        return nil
    }
    r.y0 = right.y0
    r.y1 = min(left.y1, right.y1)
    return &r
}

func (r1 *Rectangle) IsIntersect(r2 *Rectangle) bool {
    return r1.XLine().IsIntersect(r2.XLine()) &&
           r1.YLine().IsIntersect(r2.YLine())
    //return (&Line{r1.x0, r1.x1, ' '}).IsIntersect(&Line{r2.x0, r2.x1, ' '}) &&
    //       (&Line{r1.y0, r1.y1, ' '}).IsIntersect(&Line{r2.y0, r2.y1, ' '})
}

func (r *Rectangle) XLine() *Line {
    return &Line{r.x0, r.x1, ' '}
}

func (r *Rectangle) YLine() *Line {
    return &Line{r.y0, r.y1, ' '}
}

func (base *Rectangle) Difference(diff *Rectangle) []*Rectangle {
    rs := []*Rectangle{}

    if !base.IsIntersect(diff) {
        return []*Rectangle{base}
    }

    xi := base.XLine().Intersect(diff.XLine())
    yi := base.YLine().Intersect(diff.YLine())

    rs = append(rs, &Rectangle{
        base.x0, xi.x0 - 1,
        base.y0, yi.x0 - 1,
        '1',
    }, &Rectangle{
        xi.x0, xi.x1,
        base.y0, yi.x0 - 1,
        '2',
    }, &Rectangle{
        xi.x1 + 1, base.x1,
        base.y0, yi.x0 - 1,
        '3',
    }, &Rectangle{
        xi.x1 + 1, base.x1,
        yi.x0, yi.x1,
        '4',
    }, &Rectangle{
        xi.x1 + 1, base.x1,
        yi.x1 + 1, base.y1,
        '5',
    }, &Rectangle{
        xi.x0, xi.x1,
        yi.x1 + 1, base.y1,
        '6',
    }, &Rectangle{
        base.x0, xi.x0 - 1,
        yi.x1 + 1, base.y1,
        '7',
    }, &Rectangle{
        base.x0, xi.x0 - 1,
        yi.x0, yi.x1,
        '8',
    })

    valid := []*Rectangle{}
    for _, r := range rs {
        if r.IsValid() {
            valid = append(valid, r)
        }
    }

    return valid
}

