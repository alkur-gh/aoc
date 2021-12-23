package graphics

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

func isIntersectLine(ax0, ax1, bx0, bx1  int) bool {
    return ((ax0 <= bx0) && (bx0 <= ax1)) || ((bx0 <= ax0) && (ax0 <= bx1))
}

func (r1 *Rectangle) IsIntersect(r2 *Rectangle) bool {
    return isIntersectLine(r1.x0, r1.x1, r2.x0, r2.x1) &&
           isIntersectLine(r1.y0, r1.y1, r2.y0, r2.y1)
}

func (base *Rectangle) Difference(diff *Rectangle) []*Rectangle {
    rs := []*Rectangle{}

    if !base.IsIntersect(diff) {
        return []*Rectangle{base}
    }

    if (base.x0 < diff.x0) {
        rs = append(rs, &Rectangle{
            base.x0, min(diff.x0 - 1, base.x1),
            base.y0, base.y1,
            '1',
        })
    }
    if base.y0 < diff.y0 {
        rs = append(rs, &Rectangle{
            max(base.x0, diff.x0), min(diff.x1, base.x1),
            base.y0, diff.y0 - 1,
            '2',
        })
    }
    if (base.x1 > diff.x1) {
        rs = append(rs, &Rectangle{
            diff.x1 + 1, base.x1,
            base.y0, base.y1,
            '3',
        })
    }
    if (base.y1 > diff.y1) {
        rs = append(rs, &Rectangle{
            max(base.x0, diff.x0), min(diff.x1, base.x1),
            diff.y1 + 1, base.y1,
            '4',
        })
    }

    valid := []*Rectangle{}
    for _, r := range rs {
        if r.IsValid() {
            valid = append(valid, r)
        }
    }

    return valid
}

func (r1 *Rectangle) Union(r2 *Rectangle) []*Rectangle {
    rs := []*Rectangle{}
    rs = append(rs, r1.Difference(r2)...)
    rs = append(rs, r2.Difference(r1)...)
    r := r1.Intersection(r2)
    if r != nil {
        rs = append(rs, r)
    }
    if len(rs) == 0 {
        rs = append(rs, r1, r2)
    }
    return rs
}

