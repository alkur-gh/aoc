package graphics

func MakeLine(x0, x1 int, r rune) *Line {
    return &Line{x0, x1, r}
}

func (l1 *Line) IsIntersect(l2 *Line) bool {
    return ((l1.x0 <= l2.x0) && (l2.x0 <= l1.x1)) ||
           ((l2.x0 <= l1.x0) && (l1.x0 <= l2.x1))
}

func (l1 *Line) Intersect(l2 *Line) *Line {
    if !l1.IsIntersect(l2) {
        return nil
    }
    if l1.x0 <= l2.x0 {
        return &Line{l2.x0, min(l2.x1, l1.x1), 'I'}
    } else {
        return &Line{l1.x0, min(l2.x1, l1.x1), 'I'}
    }
}

func (l1 *Line) Difference(l2 *Line) []*Line {
    if !l1.IsIntersect(l2) {
        return []*Line{l1}
    }

    ls := []*Line{}

    if l1.x0 < l2.x0 {
        ls = append(ls, &Line{l1.x0, l2.x0 - 1, '1'})
    }

    if l1.x1 > l2.x1 {
        ls = append(ls, &Line{l2.x1 + 1, l1.x1, '2'})
    }

    return ls
}
